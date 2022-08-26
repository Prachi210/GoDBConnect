package helper

import (
	"log"
	"net/http"
	"sync"

	config "github.com/Prachi210/GoDBConnect/Sender/Config"
	model "github.com/Prachi210/GoDBConnect/Sender/Model"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func SendProducts(c *gin.Context) {

	log.Println("\n Request Reached to Server")

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal(configErr)
	}

	requestBody := model.Product{}
	products := []model.Product{}
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	} else {
		log.Println(requestBody)
		products = GetProducts(requestBody, config)
		log.Println(products)
	}
	if len(products) <= 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Recived The Request",
		}) // to serialize the struct into JSON and add it to the response
		log.Println(products)
		productsPerBatch := len(products) / config.CountPerRoutine
		wg := &sync.WaitGroup{}
		count := 0
		length := 0
		batchedProducts := []model.Product{}

		client := resty.New()
		for _, product := range products {
			batchedProducts = append(batchedProducts, product)
			count++
			length++
			if count == productsPerBatch {
				wg.Add(1)
				go sendData(batchedProducts, wg, config, client)
				count = 0
				batchedProducts = nil
			} else if length == len(products) && len(batchedProducts) > 0 {
				wg.Add(1)
				go sendData(batchedProducts, wg, config, client)
				batchedProducts = nil
			}
		}
		wg.Wait()

		defer sendDummyData(batchedProducts, config)
	}
}

func sendData(productBatch []model.Product, wg *sync.WaitGroup, config config.ConfigModel, client *resty.Client) {

	defer wg.Done()
	log.Println("\n ----------- BATCH: ")
	log.Println(productBatch)
	response, err := client.R().
		SetBody(productBatch).
		SetResult(&productBatch).
		Post("http://" + config.ConsumerURL + "/productsIntoDB")
	log.Println("------------------response")
	log.Println(response)
	if err != nil {
		log.Println("connecting to server..")
	}

}

//sendOEF
func sendDummyData(dummyProductSlice []model.Product, config config.ConfigModel) {

	log.Println("sending Dummy Data")
	var dummyProduct model.Product
	dummyProduct.ID = 0
	dummyProductSlice = append(dummyProductSlice, dummyProduct)
	log.Println("Dummy data is", dummyProductSlice)
	log.Println("Checking Dummy:", dummyProductSlice[0].ID)
	client := resty.New()
	response, err := client.R().
		SetBody(dummyProductSlice).
		SetResult(&dummyProductSlice).
		Post("http://" + config.ConsumerURL + "/productsIntoDB")
	log.Println("------------------ response after sending all the batches")
	log.Println(response)
	if err != nil {
		log.Println("connecting to server..")
	}
}

func CheckServiceHealth(c *gin.Context) {
	c.JSON(http.StatusOK,
		gin.H{
			"code":    http.StatusOK,
			"message": "service running",
		})
}

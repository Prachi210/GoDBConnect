package helper

import (
	"fmt"
	"log"
	"net/http"

	config "github.com/Prachi210/GoDBConnect/Consumer/Config"
	model "github.com/Prachi210/GoDBConnect/Consumer/model"
	"github.com/gin-gonic/gin"
)

func ConsumeProducts(c *gin.Context) {
	products := []model.ConsumerProduct{}
	if err := c.BindJSON(&products); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		configFile, configError := config.LoadConfig()
		if configError != nil {
			log.Fatal(err)
		}
		var rowsInserted int
		rowsInserted = ProcessData(products, configFile)
		fmt.Println("rowsInserted: ", rowsInserted)
		fmt.Println("Products:", len(products)+1)
		if len(products)+1 == rowsInserted {
			c.JSON(200, gin.H{"message": "recieved & inserted the product batches"})
			return
		} else {
			c.JSON(200, gin.H{"message": "recieving the product batches"})
			return
		}

	}
}

func CheckServiceHealth(c *gin.Context) {
	c.JSON(http.StatusOK,
		gin.H{
			"code":    http.StatusOK,
			"message": "service running",
		})
}

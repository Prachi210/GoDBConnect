package helper

import (
	"database/sql"
	"fmt"

	config "github.com/Prachi210/GoDBConnect/Sender/Config"
	model "github.com/Prachi210/GoDBConnect/Sender/Model"
	_ "github.com/lib/pq"
)

// DB set up
func setupDB(config config.ConfigModel) *sql.DB {

	dbinfo := fmt.Sprintf("host= %s port= %d dbname= %s user= %s password= %s sslmode= %s",
		config.Host, config.Port, config.Dbname, config.User, config.Password, config.Sslmode)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		fmt.Println("Err", err.Error())
		panic(err)
	}
	return db
}

func GetProducts(requestBody model.Product, config config.ConfigModel) []model.Product {
	fmt.Println("\n Setting Up Database")
	db := setupDB(config)
	defer db.Close()
	products := []model.Product{}
	queryString := "SELECT id, model, itemtype ,category,price FROM productInfo where 1=1 "
	if (requestBody != model.Product{}) {
		if requestBody.Category != "" {
			queryString = queryString + " and category ='" + requestBody.Category + "'"
		}
		if requestBody.ItemType != "" {
			queryString = queryString + " and itemtype ='" + requestBody.ItemType + "'"
		}
		if requestBody.Model != "" {
			queryString = queryString + " and model ='" + requestBody.Model + "'"
		}
	}
	fmt.Println("Query is:", queryString)
	rows, e := db.Query(queryString)
	if e != nil {
		fmt.Print(e)
	} else {
		for rows.Next() {
			var prod model.Product

			rows.Scan(&prod.ID, &prod.Model, &prod.ItemType, &prod.Category, &prod.Price)
			fmt.Printf("Product ID = %d\t, Model= %s\t, Item Type = %s ", prod.ID, prod.Model, prod.ItemType)
			products = append(products, prod)

		}
	}
	return products

}

package main

import (
	"log"

	config "github.com/Prachi210/GoDBConnect/Consumer/Config"
	helper "github.com/Prachi210/GoDBConnect/Consumer/helper"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	//Routing incoming requests
	router := gin.Default()                                //Initialize a Gin router using Default.
	router.POST("/productsIntoDB", helper.ConsumeProducts) //to associate the GET HTTP method and /products path with a handler function.
	router.GET("/", helper.CheckServiceHealth)
	router.Run(config.ConsumerURL)
}

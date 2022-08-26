package main

import (
	"log"

	config "github.com/Prachi210/GoDBConnect/Sender/Config"
	helper "github.com/Prachi210/GoDBConnect/Sender/Helper"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}
	//Routing incoming requests
	router := gin.Default()
	router.POST("/products", helper.SendProducts)
	router.GET("/check", helper.CheckServiceHealth)
	router.Run(config.SenderURL)
	//defer helper.CallLogs(infoLog)
}

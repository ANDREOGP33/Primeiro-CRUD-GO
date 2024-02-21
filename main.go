package main

import (
	"log"

	"github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/configuration/logger"
	"github.com/ANDREOGP33/Primeiro-CRUD-GO/tree/main/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("about to start application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8084"); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"github.com/farhapartex/ainventory/config"
	"github.com/farhapartex/ainventory/controller"
	"github.com/farhapartex/ainventory/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	authController := controller.NewAuthController(config.DB)

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery()) // Middleware for logging and recovery

	routes.RegisterRoute(router, authController)

	router.Run(":8000")
}

package main

import (
	"github.com/farhapartex/ainventory/config"
	"github.com/farhapartex/ainventory/controller"
	"github.com/farhapartex/ainventory/middlewares"
	"github.com/farhapartex/ainventory/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.MigrateDB()

	authController := controller.NewAuthController(config.DB)

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.CORSMiddleware())

	routes.RegisterRoute(router, authController)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8000")
}

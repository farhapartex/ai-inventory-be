package routes

import (
	"github.com/farhapartex/ainventory/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine, authController *controller.AuthController) {
	api := r.Group("/api/v1")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}

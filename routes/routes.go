package routes

import (
	"github.com/farhapartex/ainventory/controller"
	"github.com/farhapartex/ainventory/views"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine, authController *controller.AuthController) {
	api := r.Group("/api/v1")
	{
		api.POST(("/auth/signup/"), func(ctx *gin.Context) {
			views.SignUpView(ctx, authController)
		})
	}
}

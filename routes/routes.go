package routes

import (
	"github.com/farhapartex/ainventory/controller"
	"github.com/farhapartex/ainventory/middlewares"
	"github.com/farhapartex/ainventory/views"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine, authController *controller.AuthController) {
	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST(("/signup/"), func(ctx *gin.Context) {
				views.SignUpAPIView(ctx, authController)
			})
			auth.POST(("/signin/"), func(ctx *gin.Context) {
				views.SignInAPIView(ctx, authController)
			})
		}
	}

	protected := r.Group("/api/v1")
	protected.Use(middlewares.AuthMiddleware())
	{
		user := protected.Group("/user")
		{
			user.GET(("/me/"), func(ctx *gin.Context) {
				views.UserProfileAPIView(ctx, authController)
			})

			user.POST(("/onboard/"), func(ctx *gin.Context) {
				views.UserOnboardAPIView(ctx, authController)
			})
		}
		product := protected.Group("/product")
		{
			product.GET(("/categories/"), func(ctx *gin.Context) {
				views.ProductCategoryListAPIView(ctx, authController)
			})
			product.POST(("/categories/"), func(ctx *gin.Context) {
				views.ProductCategoryCreateAPIView(ctx, authController)
			})
			product.PATCH(("/categories/:id"), func(ctx *gin.Context) {
				views.ProductCategoryUpdateAPIView(ctx, authController)
			})
			product.DELETE(("/categories/:id"), func(ctx *gin.Context) {
				views.ProductCategoryDeleteAPIView(ctx, authController)
			})
			product.GET(("/suppliers/"), func(ctx *gin.Context) {
				views.SupplierListAPIView(ctx, authController)
			})
		}
	}
}

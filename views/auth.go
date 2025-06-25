package views

import (
	"net/http"

	"github.com/farhapartex/ainventory/controller"
	"github.com/farhapartex/ainventory/dto"
	"github.com/gin-gonic/gin"
)

func SignUpAPIView(ctx *gin.Context, ac *controller.AuthController) {
	var req dto.SignUpRequestDTO
	if err := ctx.ShouldBindJSON((&req)); err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid input",
		})
	}

	response, err := ac.SignUp(req)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func SignInAPIView(ctx *gin.Context, ac *controller.AuthController) {
	var req dto.SignInRequestDTO
	if err := ctx.ShouldBindJSON((&req)); err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid input",
		})
	}

	response, err := ac.SignIn(req)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

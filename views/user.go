package views

import (
	"net/http"

	"github.com/farhapartex/ainventory/controller"
	"github.com/farhapartex/ainventory/dto"
	"github.com/farhapartex/ainventory/utils"
	"github.com/gin-gonic/gin"
)

func UserProfileAPIView(ctx *gin.Context, ac *controller.AuthController) {

	user, err := utils.GetAuthenticatedUser(ctx)
	if err != nil {
		utils.HandleAuthError(ctx, err)
		return
	}

	resp, err := ac.UserProfile(user.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func UserOnboardAPIView(ctx *gin.Context, ac *controller.AuthController) {
	user, err := utils.GetAuthenticatedUser(ctx)
	if err != nil {
		utils.HandleAuthError(ctx, err)
		return
	}

	var req dto.UserOnboardRequestDTO
	if err := ctx.ShouldBindJSON((&req)); err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid input",
		})
	}

	resp, err := ac.UserOnboard(user, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

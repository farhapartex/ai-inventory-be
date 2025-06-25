package views

import (
	"net/http"

	"github.com/farhapartex/ainventory/controller"
	"github.com/farhapartex/ainventory/models"
	"github.com/gin-gonic/gin"
)

func UserProfileAPIView(ctx *gin.Context, ac *controller.AuthController) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Use not found, need user login",
		})
		return
	}
	userMode, ok := user.(models.User)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Need user login",
		})
		return
	}

	resp, err := ac.UserProfile(userMode.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

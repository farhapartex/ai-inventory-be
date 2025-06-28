package utils

import (
	"errors"
	"net/http"

	"github.com/farhapartex/ainventory/models"
	"github.com/gin-gonic/gin"
)

func GetAuthenticatedUser(ctx *gin.Context) (*models.User, error) {
	user, exists := ctx.Get("user")
	if !exists {
		return nil, errors.New("user not found, need user login")
	}

	userModel, ok := user.(models.User)
	if !ok {
		return nil, errors.New("need user login")
	}

	return &userModel, nil
}

func HandleAuthError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}

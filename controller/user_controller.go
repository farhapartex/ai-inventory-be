package controller

import (
	"errors"

	"github.com/farhapartex/ainventory/dto"
	"github.com/farhapartex/ainventory/mapper"
	"github.com/farhapartex/ainventory/models"
)

func (ac *AuthController) UserProfile(userId uint) (*dto.UserMeResponseDTO, error) {
	var user models.User

	result := ac.DB.Where("id = ?", userId).First(&user)

	if result.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	return mapper.UserModelToUserProfileDTO(&user), nil
}

package mapper

import (
	"github.com/farhapartex/ainventory/dto"
	"github.com/farhapartex/ainventory/models"
)

func UserModelToUserProfileDTO(user *models.User) *dto.UserProfileResponseDTO {
	return &dto.UserProfileResponseDTO{
		EmployeeID: user.EmployeeID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
	}
}

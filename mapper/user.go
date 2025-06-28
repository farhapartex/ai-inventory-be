package mapper

import (
	"github.com/farhapartex/ainventory/dto"
	"github.com/farhapartex/ainventory/models"
)

func UserModelToUserProfileDTO(user *models.User) *dto.UserMeResponseDTO {
	return &dto.UserMeResponseDTO{
		EmployeeID:    user.EmployeeID,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Email:         user.Email,
		Organizations: []dto.OrganizationDTO{},
	}
}

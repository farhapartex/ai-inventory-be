package mapper

import (
	"github.com/farhapartex/ainventory/dto"
	"github.com/farhapartex/ainventory/models"
)

func OrganizationModelToDTO(org *models.Organization) *dto.OrganizationDTO {
	return &dto.OrganizationDTO{
		ID:   org.ID,
		Name: org.Name,
	}
}

func UserModelToUserProfileDTO(user *models.User) *dto.UserMeResponseDTO {

	organizations := make([]dto.OrganizationDTO, len(user.Organizations))
	for i, org := range user.Organizations {
		organizations[i] = *OrganizationModelToDTO(&org)
	}

	return &dto.UserMeResponseDTO{
		EmployeeID:    user.EmployeeID,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Email:         user.Email,
		Organizations: organizations,
	}
}

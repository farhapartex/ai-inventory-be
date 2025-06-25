package mapper

import (
	"github.com/farhapartex/ainventory/dto"
	"github.com/farhapartex/ainventory/models"
)

func SignUpDTOToModel(dto dto.SignUpRequestDTO) models.User {
	return models.User{
		FirstName:     dto.FirstName,
		LastName:      dto.LastName,
		Email:         dto.Email,
		Password:      dto.Password,
		Gender:        dto.Gender,
		Status:        "active",
		IsSuperuser:   false,
		EmailVerified: true,
	}
}

package mapper

import (
	"github.com/farhapartex/ainventory/dto"
	"github.com/farhapartex/ainventory/models"
)

func SignUpDTOToModel(dto dto.SignUpRequestDTO, password string) models.User {
	return models.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Password:  password,
	}
}

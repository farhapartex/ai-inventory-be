package controller

import (
	"errors"
	"strings"

	"github.com/farhapartex/ainventory/dto"
	"github.com/farhapartex/ainventory/mapper"
	"github.com/farhapartex/ainventory/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		DB: db,
	}
}

func (ac *AuthController) SignUp(req dto.SignUpRequestDTO) (*dto.SignUpResponseDTO, error) {
	req.Email = strings.ToLower(req.Email)
	var existingUser models.User
	result := ac.DB.Where("email = ?", req.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		return nil, errors.New("email already exists")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("password hashing failed")
	}

	newUser := mapper.SignUpDTOToModel(req, string(hashPassword))

	tx := ac.DB.Begin()
	err = tx.Create(&newUser).Error
	if err != nil {
		tx.Rollback()
		return nil, errors.New("failed to create user")
	}

	// TODO: will implement email sending later

	return &dto.SignUpResponseDTO{
		IsSuccess: true,
		Message:   "User created successfully",
	}, nil

}

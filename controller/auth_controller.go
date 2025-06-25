package controller

import (
	"errors"
	"strings"

	"github.com/farhapartex/ainventory/dto"
	"github.com/farhapartex/ainventory/mapper"
	"github.com/farhapartex/ainventory/models"
	"github.com/farhapartex/ainventory/utils"
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

func (ac *AuthController) SignIn(req dto.SignInRequestDTO) (*dto.SignInResponseDTO, error) {
	req.Email = strings.ToLower((req.Email))
	var user models.User
	result := ac.DB.Where("email = ?", req.Email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("failed to retrieve user")
	}

	if user.CanLogin() != true {
		return nil, errors.New("Permission denied to login")
	}

	if user.CheckPassword(req.Password) != true {
		return nil, errors.New("invalid password")
	}

	token, err := utils.GenerateToken(user, "access")

	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &dto.SignInResponseDTO{
		Token: token,
	}, nil

}

func (ac *AuthController) SignUp(req dto.SignUpRequestDTO) (*dto.SignUpResponseDTO, error) {
	req.Email = strings.ToLower(req.Email)
	var existingUser models.User
	result := ac.DB.Where("email = ?", req.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		return nil, errors.New("email already exists")
	}

	newUser := mapper.SignUpDTOToModel(req)

	result = ac.DB.Create(&newUser)
	if result.Error != nil {
		return nil, errors.New("failed to create user")
	}

	// TODO: will implement email sending later

	return &dto.SignUpResponseDTO{
		IsSuccess: true,
		Message:   "User created successfully",
	}, nil

}

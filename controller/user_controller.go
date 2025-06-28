package controller

import (
	"errors"

	"github.com/farhapartex/ainventory/dto"
	"github.com/farhapartex/ainventory/mapper"
	"github.com/farhapartex/ainventory/models"
)

func (ac *AuthController) UserProfile(userId uint) (*dto.UserMeResponseDTO, error) {
	var user models.User

	result := ac.DB.Where("id = ?", userId).Preload("Organizations").First(&user)

	if result.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	return mapper.UserModelToUserProfileDTO(&user), nil
}

func (ac *AuthController) UserOnboard(user *models.User, req dto.UserOnboardRequestDTO) (*dto.UserOnboardResponseDTO, error) {
	tx := ac.DB.Begin()
	if tx.Error != nil {
		return nil, errors.New("failed to start transaction")
	}

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	if err := tx.Save(user).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("failed to update user")
	}

	organization := models.Organization{
		Name:    req.Organization,
		Address: req.Address,
		City:    req.City,
		State:   req.State,
		ZipCode: req.ZipCode,
		Country: req.Country,
	}
	if err := tx.Create(&organization).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("failed to create organization")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, errors.New("Failed to save data")
	}

	return &dto.UserOnboardResponseDTO{
		OrganizationID: organization.ID,
		Organization:   organization.Name,
	}, nil
}

package controller

import (
	"errors"

	"github.com/farhapartex/ainventory/dto"
	"github.com/farhapartex/ainventory/models"
)

func (ac *AuthController) ProductCategoryList(page, pageSize int) (*[]dto.ProductCategoryResponseDTO, error) {
	var categories []models.ProductCategory

	offset := (page - 1) * pageSize
	query := ac.DB.Model(&models.ProductCategory{})
	err := query.Offset(offset).Limit(pageSize).Find(&categories).Error
	if err != nil {
		return nil, errors.New("error retrieving properties")
	}

	var responseDTOs []dto.ProductCategoryResponseDTO
	for _, category := range categories {
		dto := dto.ProductCategoryResponseDTO{
			ID:           category.ID,
			Name:         category.Name,
			ParentID:     category.ParentID,
			SortOrder:    category.SortOrder,
			IsActive:     category.IsActive,
			CreatedAt:    category.CreatedAt,
			UpdatedAt:    category.UpdatedAt,
			ProductCount: 0,
		}
		responseDTOs = append(responseDTOs, dto)
	}

	return &responseDTOs, nil
}

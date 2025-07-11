package controller

import (
	"errors"

	"github.com/farhapartex/ainventory/dto"
	"github.com/farhapartex/ainventory/mapper"
	"github.com/farhapartex/ainventory/models"
)

func (ac *AuthController) ProductCategoryList(page, pageSize int) (*dto.PaginatedResponse, error) {
	var categories []models.ProductCategory
	var totalCount int64

	ac.DB.Model(&models.ProductCategory{}).Count(&totalCount)

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
			Code:         category.Code,
			Description:  category.Description,
			ParentID:     category.ParentID,
			SortOrder:    category.SortOrder,
			IsActive:     category.IsActive,
			CreatedAt:    category.CreatedAt,
			UpdatedAt:    category.UpdatedAt,
			ProductCount: 0,
		}
		responseDTOs = append(responseDTOs, dto)
	}

	return &dto.PaginatedResponse{
		Data:       responseDTOs,
		TotalPages: totalCount,
		Page:       page,
		PageSize:   pageSize,
		Total:      len(responseDTOs),
	}, nil

	//return &responseDTOs, nil
}

func (ac *AuthController) CreateProductCategoryController(request dto.ProductCategoryRequestDTO) (*dto.ProductCategoryResponseDTO, error) {
	var category models.ProductCategory

	result := ac.DB.Where("code = ?", request.Code).First(&category)
	if result.RowsAffected > 0 {
		return nil, errors.New("Category exists")
	}

	if request.ParentID != nil {
		var parentCategory models.ProductCategory
		result = ac.DB.Where("id = ?", request.ParentID).First(&parentCategory)
		if result.RowsAffected == 0 {
			return nil, errors.New("Parent category doesn not exists")
		}
	}

	err := request.Validate()
	if err != nil {
		return nil, err
	}
	request.Normalize()

	newRow := mapper.ProductCategoryDTOToModel(request)
	result = ac.DB.Create(newRow)
	if result.Error != nil {
		return nil, errors.New("failed to create category")
	}

	return mapper.ProductCategoryModelToDTO(*newRow), nil
}

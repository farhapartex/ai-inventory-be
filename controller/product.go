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

func (ac *AuthController) UpdateProductCategoryController(
	categoryID uint,
	request dto.ProductCategoryRequestDTO,
) (*dto.ProductCategoryResponseDTO, error) {
	var category models.ProductCategory

	result := ac.DB.Where("id = ?", categoryID).First(&category)
	if result.RowsAffected == 0 {
		return nil, errors.New("category not found")
	}

	var existingCategory models.ProductCategory
	result = ac.DB.Where("code = ? AND id != ?", request.Code, categoryID).First(&existingCategory)
	if result.RowsAffected > 0 {
		return nil, errors.New("category code already exists")
	}

	if request.ParentID != nil {
		var parentCategory models.ProductCategory
		result = ac.DB.Where("id = ?", request.ParentID).First(&parentCategory)
		if result.RowsAffected == 0 {
			return nil, errors.New("parent category does not exist")
		}

		if *request.ParentID == categoryID {
			return nil, errors.New("category cannot be its own parent")
		}
	}

	err := request.Validate()
	if err != nil {
		return nil, err
	}
	request.Normalize()

	updateData := mapper.ProductCategoryDTOToModel(request)
	updateData.ID = categoryID

	result = ac.DB.Model(&category).Updates(updateData)
	if result.Error != nil {
		return nil, errors.New("failed to update category")
	}

	result = ac.DB.Where("id = ?", categoryID).First(&category)
	if result.Error != nil {
		return nil, errors.New("failed to fetch updated category")
	}

	return mapper.ProductCategoryModelToDTO(category), nil
}

func (ac *AuthController) DeleteProductCategoryController(categoryID uint) error {
	var category models.ProductCategory

	result := ac.DB.Where("id = ?", categoryID).First(&category)
	if result.RowsAffected == 0 {
		return errors.New("category not found")
	}

	var productCount int64
	result = ac.DB.Model(&models.Product{}).Where("category_id = ?", categoryID).Count(&productCount)
	if result.Error != nil {
		return errors.New("Failed to check products")
	}
	if productCount > 0 {
		return errors.New("Cannot delete category with associated products")
	}

	var childCount int64
	result = ac.DB.Model(&models.ProductCategory{}).Where("parent_id = ?", categoryID).Count(&childCount)
	if result.Error != nil {
		return errors.New("Failed to check child categories")
	}
	if childCount > 0 {
		return errors.New("Cannot delete category with child categories")
	}

	result = ac.DB.Delete(&category)
	if result.Error != nil {
		return errors.New("Failed to delete category")
	}

	return nil
}

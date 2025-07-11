package mapper

import (
	"github.com/farhapartex/ainventory/dto"
	"github.com/farhapartex/ainventory/models"
)

func ProductCategoryDTOToModel(data dto.ProductCategoryRequestDTO) *models.ProductCategory {
	model := models.ProductCategory{
		Name:        data.Name,
		Code:        data.Code,
		Description: data.Description,
		ParentID:    data.ParentID,
		SortOrder:   data.SortOrder,
	}

	if data.IsActive != nil {
		model.IsActive = *data.IsActive
	} else {
		model.IsActive = true
	}

	return &model
}

func ProductCategoryModelToDTO(data models.ProductCategory) *dto.ProductCategoryResponseDTO {
	return &dto.ProductCategoryResponseDTO{
		ID:          data.ID,
		Name:        data.Name,
		Code:        data.Code,
		Description: data.Description,
		ParentID:    data.ParentID,
		SortOrder:   data.SortOrder,
		IsActive:    data.IsActive,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

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

func SupplierModelToDTO(supplier models.Supplier, productCount int) *dto.SupplierResponseDTO {
	dto := dto.SupplierResponseDTO{
		ID:                supplier.ID,
		Name:              supplier.Name,
		Code:              supplier.Code,
		ContactPerson:     supplier.ContactPerson,
		Email:             supplier.Email,
		Phone:             supplier.Phone,
		Website:           supplier.Website,
		Address:           supplier.Address,
		City:              supplier.City,
		State:             supplier.State,
		ZipCode:           supplier.ZipCode,
		Country:           supplier.Country,
		TaxID:             supplier.TaxID,
		PaymentTerms:      supplier.PaymentTerms,
		Currency:          supplier.Currency,
		MinimumOrderValue: supplier.MinimumOrderValue,
		Status:            supplier.Status,
		Rating:            supplier.Rating,
		Notes:             supplier.Notes,
		CreatedBy:         supplier.CreatedBy,
		CreatedAt:         supplier.CreatedAt,
		UpdatedAt:         supplier.UpdatedAt,
		ProductCount:      productCount,
	}

	return &dto
}

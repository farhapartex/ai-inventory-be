package dto

import (
	"fmt"
	"strings"
	"time"
)

type ProductCategoryRequestDTO struct {
	Name        string `json:"name" binding:"required,min=2,max=100" validate:"required,min=2,max=100"`
	Code        string `json:"code" binding:"required,min=2,max=20,alphanum_underscore" validate:"required,min=2,max=20"`
	Description string `json:"description" binding:"omitempty,max=500" validate:"omitempty,max=500"`
	ParentID    *uint  `json:"parent_id" binding:"omitempty,min=1" validate:"omitempty,min=1"`
	SortOrder   int    `json:"sort_order" binding:"omitempty,min=0,max=9999" validate:"omitempty,min=0,max=9999"`
	IsActive    *bool  `json:"is_active" binding:"omitempty" validate:"omitempty"`
}

type ProductCategoryResponseDTO struct {
	ID           uint                         `json:"id"`
	Name         string                       `json:"name"`
	Code         string                       `json:"code"`
	Description  string                       `json:"description"`
	ParentID     *uint                        `json:"parent_id"`
	SortOrder    int                          `json:"sort_order"`
	IsActive     bool                         `json:"is_active"`
	CreatedAt    time.Time                    `json:"created_at"`
	UpdatedAt    time.Time                    `json:"updated_at"`
	Parent       *ProductCategoryResponseDTO  `json:"parent,omitempty"`
	Children     []ProductCategoryResponseDTO `json:"children,omitempty"`
	ProductCount int                          `json:"product_count,omitempty"`
}

func (dto *ProductCategoryRequestDTO) Normalize() {
	dto.Name = strings.TrimSpace(dto.Name)
	dto.Code = strings.ToUpper(strings.TrimSpace(dto.Code))
	dto.Description = strings.TrimSpace(dto.Description)

	if dto.SortOrder == 0 {
		dto.SortOrder = 0
	}

	if dto.IsActive == nil {
		active := true
		dto.IsActive = &active
	}
}

func (dto *ProductCategoryRequestDTO) Validate() error {
	if dto.Code != "" {
		for _, char := range dto.Code {
			if !((char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') || char == '_') {
				return fmt.Errorf("code must contain only uppercase letters, numbers, and underscores")
			}
		}
	}

	return nil
}

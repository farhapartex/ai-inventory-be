package models

import "time"

type Organization struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null;unique"`
	Address   string    `json:"address" gorm:"size:500"`
	City      string    `json:"city" gorm:"size:100"`
	State     string    `json:"state" gorm:"size:100"`
	ZipCode   string    `json:"zip_code" gorm:"size:20"`
	Country   string    `json:"country" gorm:"size:100;default:'United States'"`
	OwnerID   uint      `json:"owner_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Owner User `json:"owner" gorm:"foreignKey:OwnerID;references:ID"`
}

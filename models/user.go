package models

import "time"

type User struct {
	ID                    uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName             string     `json:"first_name" gorm:"size:150;not null"`
	LastName              string     `json:"last_name" gorm:"size:150;not null"`
	Email                 string     `json:"email" gorm:"size:150;not null;unique"`
	Password              string     `json:"password" gorm:"size:255;not null"`
	IsSuperuser           bool       `gorm:"default:false" json:"is_superuser"`
	JoinedAt              time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"joined_at"`
	LastLoginAt           *time.Time `json:"last_login_at"`
	LastPasswordChangedAt *time.Time `json:"password_changed_at"`
	Status                string     `gorm:"size:20;default:active;check:status IN ('active', 'inactive', 'suspended')" json:"status"`
	EmailVerified         bool       `gorm:"default:false" json:"email_verified"`
	VerifiedAt            *time.Time `json:"verified_at"`
}

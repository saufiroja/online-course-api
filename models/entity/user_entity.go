package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              int64          `json:"id"`
	Name            string         `json:"name"`
	Email           string         `json:"email"`
	Password        string         `json:"-"`
	CodeVerified    string         `json:"-"`
	EmailVerifiedAt *time.Time     `json:"email_verified_at"`
	CreatedByID     *int64         `json:"created_by" gorm:"column:created_by"`
	CreatedBy       *Admin         `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID     *int64         `json:"updated_by" gorm:"column:updated_by"`
	UpdatedBy       *Admin         `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt       *time.Time     `json:"created_at"`
	UpdatedAt       *time.Time     `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}

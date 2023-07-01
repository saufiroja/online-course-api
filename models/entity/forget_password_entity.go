package entity

import (
	"time"

	"gorm.io/gorm"
)

type ForgotPassword struct {
	ID        int64          `json:"id"`
	User      *User          `json:"user" gorm:"foreignKey:UserID;references:ID"`
	UserID    *int64         `json:"user_id"`
	Valid     bool           `json:"valid"`
	Code      string         `json:"code"`
	ExpiredAt *time.Time     `json:"expired_at"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

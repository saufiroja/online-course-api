package entity

import (
	"time"

	"gorm.io/gorm"
)

type ClassRoom struct {
	ID          int64          `json:"id"`
	User        *User          `json:"user" gorm:"foreignKey:UserID;references:ID"`
	UserID      int64          `json:"user_id"`
	Product     *Product       `json:"product" gorm:"foreignKey:ProductID;references:ID"`
	ProductID   *int64         `json:"product_id"`
	CreatedBy   *User          `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	CreatedByID *int64         `json:"created_by" gorm:"column:created_by"`
	UpdateBy    *User          `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID *int64         `json:"updated_by" gorm:"column:updated_by"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

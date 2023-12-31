package entity

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID          int64          `json:"id"`
	User        *User          `json:"user" gorm:"foreignKey:UserID;references:ID"`
	UserID      *int64         `json:"user_id"`
	Product     *Product       `json:"product" gorm:"foreingKey:ProductID;references:ID"`
	ProductID   *int64         `json:"product_id"`
	Quantity    int64          `json:"quantity"`
	IsChecked   bool           `json:"is_checked"`
	CreatedBy   *User          `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	CreatedByID *int64         `json:"created_by" gorm:"column:created_by"`
	UpdatedByID *int64         `json:"updated_by" gorm:"column:updated_by"`
	UpdatedBy   *User          `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

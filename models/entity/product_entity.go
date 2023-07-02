package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID                int64            `json:"id"`
	ProductCategoryID *int64           `json:"product_category_id"`
	ProductCategory   *ProductCategory `json:"product_category" gorm:"foreingKey:ProductCategoryID;references:ID"`
	Title             string           `json:"title"`
	Image             *string          `json:"image"`
	Video             *string          `json:"-"`
	VideoLink         *string          `json:"video,omitempty" gorm:"-"`
	Description       string           `json:"description"`
	IsHighlighted     bool             `json:"is_highlighted"`
	Price             int64            `json:"price"`
	CreatedByID       *int64           `json:"created_by" gorm:"column:created_by"`
	CreatedBy         *Admin           `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID       *int64           `json:"updated_by" gorm:"column:updated_by"`
	UpdatedBy         *Admin           `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt         *time.Time       `json:"created_at"`
	UpdatedAt         *time.Time       `json:"updated_at"`
	DeletedAt         gorm.DeletedAt   `json:"deleted_at"`
}

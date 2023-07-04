package entity

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID           int64          `json:"id"`
	User         *User          `json:"user" gorm:"foreignKey:UserID;references:ID"`
	UserID       *int64         `json:"user_id"`
	OrderDetails []OrderDetail  `json:"order_details"`
	Discount     *Discount      `json:"discount" gorm:"foreignKey:DiscountID;references:ID"`
	DiscountID   *int64         `json:"discount_id"`
	CheckoutLink string         `json:"checkout_link"`
	ExternalID   string         `json:"external_id"`
	Price        int64          `json:"price"`
	TotalPrice   int64          `json:"total_price"`
	Status       string         `json:"status"`
	CreatedBy    *User          `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	CreatedByID  *int64         `json:"created_by" gorm:"column:created_by"`
	UpdatedByID  *int64         `json:"updated_by" gorm:"column:updated_by"`
	UpdatedBy    *User          `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt    *time.Time     `json:"created_at"`
	UpdatedAt    *time.Time     `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

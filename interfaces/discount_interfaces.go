package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/models/entity"
)

type DiscountRepository interface {
	FindAllDiscount(offset, limit int) ([]entity.Discount, error)
	FindDiscountById(id int) (*entity.Discount, error)
	FindDiscountByCode(code string) (*entity.Discount, error)
	InsertDiscount(input entity.Discount) (*entity.Discount, error)
	UpdateDiscountById(input entity.Discount) (*entity.Discount, error)
	DeleteDiscountById(input entity.Discount) error
}

type DiscountService interface {
	FindAllDiscount(offset, limit int) ([]entity.Discount, error)
	FindDiscountById(id int) (*entity.Discount, error)
	FindDiscountByCode(code string) (*entity.Discount, error)
	InsertDiscount(input dto.DiscountRequestBody) (*entity.Discount, error)
	UpdateDiscountById(id int, input dto.DiscountRequestBody) (*entity.Discount, error)
	DeleteDiscountById(id int) error
	UpdateRemainingDiscount(id int, input dto.DiscountRemainingQuantityRequestBody) (*entity.Discount, error)
}

type DiscountHandler interface {
	FindAllDiscount(c *fiber.Ctx) error
	FindDiscountById(c *fiber.Ctx) error
	FindDiscountByCode(c *fiber.Ctx) error
	InsertDiscount(c *fiber.Ctx) error
	UpdateDiscountById(c *fiber.Ctx) error
	DeleteDiscountById(c *fiber.Ctx) error
	UpdateRemainingDiscount(c *fiber.Ctx) error
}

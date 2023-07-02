package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/models/entity"
)

type ProductCategoryRepository interface {
	FindAllProductCategory(offset, limit int) ([]entity.ProductCategory, error)
	FindProductCategoryByID(id int) (entity.ProductCategory, error)
	InsertProductCategory(input entity.ProductCategory) (entity.ProductCategory, error)
	UpdateProductCategory(input entity.ProductCategory) (entity.ProductCategory, error)
	DeleteProductCategory(input entity.ProductCategory) error
}

type ProductCategoryService interface {
	FindAllProductCategory(offset, limit int) ([]entity.ProductCategory, error)
	FindProductCategoryByID(id int) (entity.ProductCategory, error)
	InsertProductCategory(input dto.ProductCategoryRequestBody) (entity.ProductCategory, error)
	UpdateProductCategory(id int, input dto.ProductCategoryRequestBody) (entity.ProductCategory, error)
	DeleteProductCategory(id int) error
}

type ProductCategoryHandler interface {
	FindAllProductCategory(c *fiber.Ctx) error
	FindProductCategoryByID(c *fiber.Ctx) error
	InsertProductCategory(c *fiber.Ctx) error
	UpdateProductCategory(c *fiber.Ctx) error
	DeleteProductCategory(c *fiber.Ctx) error
}

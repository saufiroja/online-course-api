package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/models/entity"
)

type ProductRepository interface {
	FindAllProduct(offset, limit int) ([]entity.Product, error)
	FindProductByID(id int) (*entity.Product, error)
	InsertProduct(input entity.Product) (*entity.Product, error)
	UpdateProductByID(input entity.Product) (*entity.Product, error)
	DeleteProductByID(input entity.Product) error
	CountProduct() (int64, error)
}

type ProductServices interface {
	FindAllProduct(offset, limit int) ([]entity.Product, error)
	FindProductByID(id int) (*entity.Product, error)
	InsertProduct(input dto.ProductRequestBody) (*entity.Product, error)
	UpdateProductByID(id int, input dto.ProductRequestBody) (*entity.Product, error)
	DeleteProductByID(id int) error
	CountProduct() (int64, error)
}

type ProductHandler interface {
	FindAllProduct(c *fiber.Ctx) error
	FindProductByID(c *fiber.Ctx) error
	InsertProduct(c *fiber.Ctx) error
	UpdateProductByID(c *fiber.Ctx) error
	DeleteProductByID(c *fiber.Ctx) error
}

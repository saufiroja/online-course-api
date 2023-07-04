package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/models/entity"
)

type CartRepository interface {
	FindAllCartByUserId(userId, offset, limit int) ([]entity.Cart, error)
	FindCartById(id int) (*entity.Cart, error)
	InsertCart(input entity.Cart) (*entity.Cart, error)
	UpdateCart(input entity.Cart) (*entity.Cart, error)
	DeleteCart(input entity.Cart) error
	DeleteCartByUserId(userId int) error
}

type CartService interface {
	FindAllCartByUserId(userId, offset, limit int) ([]entity.Cart, error)
	FindCartById(id int) (*entity.Cart, error)
	InsertCart(input dto.CartRequestBody) (*entity.Cart, error)
	UpdateCart(id int, input dto.CartRequestUpdateBody) (*entity.Cart, error)
	DeleteCart(id, userId int) error
	DeleteCartByUserId(userId int) error
}

type CartHandler interface {
	FindAllCartByUserId(c *fiber.Ctx) error
	InsertCart(c *fiber.Ctx) error
	DeleteCart(c *fiber.Ctx) error
	UpdateCart(c *fiber.Ctx) error
}

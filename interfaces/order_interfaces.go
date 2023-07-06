package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/models/entity"
)

type OrderRepository interface {
	FindAllOrdersByUserId(userId, offset, limit int) ([]entity.Order, error)
	FindOrderByExternalId(externalId string) (entity.Order, error)
	FindOrderById(id int) (entity.Order, error)
	InsertOrder(input entity.Order) (entity.Order, error)
	UpdateOrder(input entity.Order) (entity.Order, error)
	CountOrder() (int64, error)
}

type OrderService interface {
	FindAllOrdersByUserId(userId, offset, limit int) ([]entity.Order, error)
	FindOrderByExternalId(externalId string) (*entity.Order, error)
	FindOrderById(id int) (entity.Order, error)
	InsertOrder(input dto.OrderRequestBody) (entity.Order, error)
	UpdateOrder(id int, input dto.OrderRequestBody) (entity.Order, error)
	CountOrder() (int64, error)
}

type OrderHandler interface {
	FindAllOrdersByUserId(c *fiber.Ctx) error
	FindOrderById(c *fiber.Ctx) error
	InsertOrder(c *fiber.Ctx) error
}

package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/models/dto"
)

type RegisterService interface {
	Register(input *dto.UserRequestBody) error
}

type RegisterHandler interface {
	Register(c *fiber.Ctx) error
}

package register

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/utils"
)

type handler struct {
	s interfaces.RegisterService
}

func NewRegisterHandler(s interfaces.RegisterService) interfaces.RegisterHandler {
	return &handler{s}
}

// Register implements interfaces.RegisterHandler.
func (h *handler) Register(c *fiber.Ctx) error {
	var input dto.UserRequestBody
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	err := h.s.Register(&input)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"message": "success registered, please check your email to verify your account",
	})
}

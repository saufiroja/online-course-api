package forgetPassword

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/utils"
)

type handler struct {
	s interfaces.ForgetPasswordService
}

func NewForgetPasswordHandler(s interfaces.ForgetPasswordService) interfaces.ForgetPasswordHandler {
	return &handler{
		s: s,
	}
}

// InsertForgetPassword implements interfaces.ForgetPasswordHandler.
func (h *handler) InsertForgetPassword(c *fiber.Ctx) error {
	var input dto.ForgotPasswordRequestBody

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    fiber.StatusBadRequest,
		})
	}

	_, err := h.s.InsertForgetPassword(&input)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success, please check your email",
		"code":    fiber.StatusCreated,
	})
}

// UpdateForgetPassword implements interfaces.ForgetPasswordHandler.
func (h *handler) UpdateForgetPassword(c *fiber.Ctx) error {
	var input dto.ForgotPasswordUpdateRequestBody

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    fiber.StatusBadRequest,
		})
	}

	_, err := h.s.UpdateForgetPassword(&input)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success change password",
		"code":    fiber.StatusOK,
	})
}

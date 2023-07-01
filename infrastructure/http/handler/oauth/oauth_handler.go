package oauth

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/utils"
)

type handler struct {
	s interfaces.OauthService
}

func NewOauthHandler(s interfaces.OauthService) interfaces.OauthHandler {
	return &handler{s}
}

// Login implements interfaces.OauthHandler.
func (h *handler) Login(c *fiber.Ctx) error {
	var input dto.LoginRequestBody
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
			"code":    400,
		})
	}

	res, err := h.s.Login(&input)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success login",
		"code":    200,
		"result":  res,
	})
}

// RefreshToken implements interfaces.OauthHandler.
func (h *handler) RefreshToken(c *fiber.Ctx) error {
	var input dto.RefreshTokenRequestBody

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
			"code":    400,
		})
	}

	res, err := h.s.RefreshToken(&input)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success refresh token",
		"code":    200,
		"result":  res,
	})
}

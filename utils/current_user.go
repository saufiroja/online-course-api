package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/models/dto"
)

func CurrentUser(c *fiber.Ctx) *dto.MapClaimsResponse {
	user := c.Locals("user")
	users := user.(*dto.MapClaimsResponse)

	return users
}

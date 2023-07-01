package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/models/dto"
)

func MiddlewaresAdmin(c *fiber.Ctx) error {
	user := c.Locals("user")

	users := user.(*dto.MapClaimsResponse)

	if !users.IsAdmin {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "you are not admin",
			"code":    401,
		})
	}

	return c.Next()
}

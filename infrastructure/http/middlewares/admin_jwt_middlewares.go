package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/utils"
)

func MiddlewaresAdmin(c *fiber.Ctx) error {
	user := utils.CurrentUser(c)

	if !user.IsAdmin {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "you are not admin",
			"code":    401,
		})
	}

	return c.Next()
}

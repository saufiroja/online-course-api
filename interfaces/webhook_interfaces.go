package interfaces

import "github.com/gofiber/fiber/v2"

type WebhookService interface {
	UpdatePayment(id string) error
}

type WebhookHandler interface {
	UpdatePayment(c *fiber.Ctx) error
}

package webhook

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/utils"
)

type handler struct {
	s interfaces.WebhookService
}

func NewWebhookHandler(s interfaces.WebhookService) interfaces.WebhookHandler {
	return &handler{s}
}

// UpdatePayment implements interfaces.WebhookHandler.
func (h *handler) UpdatePayment(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.s.UpdatePayment(id)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success update payment",
		"code":    200,
	})
}

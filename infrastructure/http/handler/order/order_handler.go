package order

import (
	"strconv"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/utils"
)

type handler struct {
	s interfaces.OrderService
}

func NewOrderHandler(s interfaces.OrderService) interfaces.OrderHandler {
	return &handler{s}
}

// FindAllOrdersByUserId implements interfaces.OrderHandler.
func (h *handler) FindAllOrdersByUserId(c *fiber.Ctx) error {
	offset := c.Query("offset")
	limit := c.Query("limit")

	offsetInt, _ := strconv.Atoi(offset)
	limitInt, _ := strconv.Atoi(limit)

	user := utils.CurrentUser(c)

	orders, err := h.s.FindAllOrdersByUserId(int(user.ID), offsetInt, limitInt)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success get orders",
		"code":    fiber.StatusOK,
		"result":  orders,
	})
}

// FindOrderById implements interfaces.OrderHandler.
func (h *handler) FindOrderById(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, _ := strconv.Atoi(id)

	order, err := h.s.FindOrderById(idInt)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success get order",
		"code":    fiber.StatusOK,
		"result":  order,
	})
}

// InsertOrder implements interfaces.OrderHandler.
func (h *handler) InsertOrder(c *fiber.Ctx) error {
	var input dto.OrderRequestBody

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error bad request",
			"code":    fiber.StatusBadRequest,
		})
	}

	user := utils.CurrentUser(c)

	input.UserID = user.ID
	input.Email = user.Email

	order, err := h.s.InsertOrder(input)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success insert order",
		"code":    fiber.StatusOK,
		"result":  order,
	})
}

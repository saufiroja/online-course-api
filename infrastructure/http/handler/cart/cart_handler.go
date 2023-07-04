package cart

import (
	"strconv"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/utils"
)

type handler struct {
	s interfaces.CartService
}

func NewCartHandler(s interfaces.CartService) interfaces.CartHandler {
	return &handler{
		s: s,
	}
}

// DeleteCart implements interfaces.CartHandler.
func (h *handler) DeleteCart(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, _ := strconv.Atoi(id)

	user := utils.CurrentUser(c)

	err := h.s.DeleteCart(idInt, int(user.ID))
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success delete cart",
		"code":    fiber.StatusOK,
	})
}

// FindAllCartByUserId implements interfaces.CartHandler.
func (h *handler) FindAllCartByUserId(c *fiber.Ctx) error {
	offset := c.Query("offset")
	limit := c.Query("limit")

	offsetInt, _ := strconv.Atoi(offset)
	limitInt, _ := strconv.Atoi(limit)

	user := utils.CurrentUser(c)

	res, err := h.s.FindAllCartByUserId(int(user.ID), offsetInt, limitInt)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success get all cart",
		"code":    fiber.StatusOK,
		"result":  res,
	})
}

// InsertCart implements interfaces.CartHandler.
func (h *handler) InsertCart(c *fiber.Ctx) error {
	var input dto.CartRequestBody

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    fiber.StatusBadRequest,
		})
	}

	user := utils.CurrentUser(c)
	input.UserID = user.ID
	input.CreatedBy = input.UserID

	res, err := h.s.InsertCart(input)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success insert cart",
		"code":    fiber.StatusCreated,
		"result":  res,
	})
}

// UpdateCart implements interfaces.CartHandler.
func (h *handler) UpdateCart(c *fiber.Ctx) error {
	var input dto.CartRequestUpdateBody

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    fiber.StatusBadRequest,
		})
	}

	id := c.Params("id")

	idInt, _ := strconv.Atoi(id)

	user := utils.CurrentUser(c)

	input.UserID = &user.ID

	res, err := h.s.UpdateCart(idInt, input)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success update cart",
		"code":    fiber.StatusOK,
		"result":  res,
	})
}

package discount

import (
	"strconv"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/utils"
)

type handler struct {
	s interfaces.DiscountService
}

func NewDiscountHandler(s interfaces.DiscountService) interfaces.DiscountHandler {
	return &handler{s}
}

// DeleteDiscountById implements interfaces.DiscountHandler.
func (h *handler) DeleteDiscountById(c *fiber.Ctx) error {
	id := c.Params("id")

	ids, _ := strconv.Atoi(id)

	err := h.s.DeleteDiscountById(ids)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success delete discount",
		"code":    200,
	})
}

// FindAllDiscount implements interfaces.DiscountHandler.
func (h *handler) FindAllDiscount(c *fiber.Ctx) error {
	offset := c.Query("offset")
	limit := c.Query("limit")

	offsets, _ := strconv.Atoi(offset)
	limits, _ := strconv.Atoi(limit)

	discount, err := h.s.FindAllDiscount(offsets, limits)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success get discounts",
		"code":    200,
		"result":  discount,
	})
}

// FindDiscountByCode implements interfaces.DiscountHandler.
func (h *handler) FindDiscountByCode(c *fiber.Ctx) error {
	code := c.Params("code")

	discount, err := h.s.FindDiscountByCode(code)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success get discount by code",
		"code":    200,
		"result":  discount,
	})
}

// FindDiscountById implements interfaces.DiscountHandler.
func (h *handler) FindDiscountById(c *fiber.Ctx) error {
	id := c.Params("id")

	ids, _ := strconv.Atoi(id)

	discount, err := h.s.FindDiscountById(ids)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success get discount by id",
		"code":    200,
		"result":  discount,
	})
}

// InsertDiscount implements interfaces.DiscountHandler.
func (h *handler) InsertDiscount(c *fiber.Ctx) error {
	var input dto.DiscountRequestBody

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "error parsing request body",
			"code":    400,
		})
	}

	admin := utils.CurrentUser(c)
	input.CreatedBy = &admin.ID

	discount, err := h.s.InsertDiscount(input)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "success insert discount",
		"code":    201,
		"result":  discount,
	})
}

// UpdateDiscountById implements interfaces.DiscountHandler.
func (h *handler) UpdateDiscountById(c *fiber.Ctx) error {
	id := c.Params("id")

	ids, _ := strconv.Atoi(id)

	var input dto.DiscountRequestBody

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "error parsing request body",
			"code":    400,
		})
	}

	admin := utils.CurrentUser(c)
	input.UpdatedBy = &admin.ID

	discount, err := h.s.UpdateDiscountById(ids, input)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success update discount",
		"code":    200,
		"result":  discount,
	})
}

// UpdateRemainingDiscount implements interfaces.DiscountHandler.
func (h *handler) UpdateRemainingDiscount(c *fiber.Ctx) error {
	code := c.Params("id")

	ids, _ := strconv.Atoi(code)

	var input dto.DiscountRemainingQuantityRequestBody

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "error parsing request body",
			"code":    400,
		})
	}

	discount, err := h.s.UpdateRemainingDiscount(ids, input)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success update discount",
		"code":    200,
		"result":  discount,
	})
}

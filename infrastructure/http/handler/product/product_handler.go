package product

import (
	"strconv"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/utils"
)

type handler struct {
	s interfaces.ProductServices
}

func NewProductHandler(s interfaces.ProductServices) interfaces.ProductHandler {
	return &handler{
		s,
	}
}

// DeleteProductByID implements interfaces.ProductHandler.
func (h *handler) DeleteProductByID(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, _ := strconv.Atoi(id)

	err := h.s.DeleteProductByID(idInt)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success delete product",
		"code":    fiber.StatusOK,
	})
}

// FindAllProduct implements interfaces.ProductHandler.
func (h *handler) FindAllProduct(c *fiber.Ctx) error {
	offset := c.Query("offset")
	limit := c.Query("limit")

	offsetInt, _ := strconv.Atoi(offset)
	limitInt, _ := strconv.Atoi(limit)

	product, err := h.s.FindAllProduct(offsetInt, limitInt)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success get all product",
		"code":    fiber.StatusOK,
		"result":  product,
	})
}

// FindProductByID implements interfaces.ProductHandler.
func (h *handler) FindProductByID(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, _ := strconv.Atoi(id)

	product, err := h.s.FindProductByID(idInt)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success get product by id",
		"code":    fiber.StatusOK,
		"result":  product,
	})
}

// InsertProduct implements interfaces.ProductHandler.
func (h *handler) InsertProduct(c *fiber.Ctx) error {
	image, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    fiber.StatusBadRequest,
		})
	}

	video, _ := c.FormFile("video")

	var input dto.ProductRequestBody

	input.Image = image
	input.Video = video

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    fiber.StatusBadRequest,
		})
	}

	admin := utils.CurrentUser(c)
	input.CreatedBy = &admin.ID

	product, err := h.s.InsertProduct(input)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success insert product",
		"code":    fiber.StatusOK,
		"result":  product,
	})
}

// UpdateProductByID implements interfaces.ProductHandler.
func (h *handler) UpdateProductByID(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, _ := strconv.Atoi(id)

	var input dto.ProductRequestBody

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    fiber.StatusBadRequest,
		})
	}

	admin := utils.CurrentUser(c)
	input.UpdatedBy = &admin.ID

	product, err := h.s.UpdateProductByID(idInt, input)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success update product",
		"code":    fiber.StatusOK,
		"result":  product,
	})
}

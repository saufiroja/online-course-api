package productcategory

import (
	"fmt"
	"strconv"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/utils"
)

type handler struct {
	s interfaces.ProductCategoryService
}

func NewProductCategoryHandler(s interfaces.ProductCategoryService) interfaces.ProductCategoryHandler {
	return &handler{s}
}

// DeleteProductCategory implements interfaces.ProductCategoryHandler.
func (h *handler) DeleteProductCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, _ := strconv.Atoi(id)

	err := h.s.DeleteProductCategory(idInt)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success delete product category",
		"code":    200,
	})
}

// FindAllProductCategory implements interfaces.ProductCategoryHandler.
func (h *handler) FindAllProductCategory(c *fiber.Ctx) error {
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	productCategories, err := h.s.FindAllProductCategory(offset, limit)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success get product categories",
		"code":    200,
		"result":  productCategories,
	})
}

// FindProductCategoryByID implements interfaces.ProductCategoryHandler.
func (h *handler) FindProductCategoryByID(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, _ := strconv.Atoi(id)

	productCategory, err := h.s.FindProductCategoryByID(idInt)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success get product category",
		"code":    200,
		"result":  productCategory,
	})
}

// InsertProductCategory implements interfaces.ProductCategoryHandler.
func (h *handler) InsertProductCategory(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": fmt.Sprintf("error when upload image: %s", err.Error()),
			"code":    400,
		})
	}

	var input dto.ProductCategoryRequestBody
	input.Image = file

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
			"code":    400,
		})
	}

	admin := utils.CurrentUser(c)
	input.CreatedBy = &admin.ID

	productCategory, err := h.s.InsertProductCategory(input)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "success insert product category",
		"code":    201,
		"result":  productCategory,
	})
}

// UpdateProductCategory implements interfaces.ProductCategoryHandler.
func (h *handler) UpdateProductCategory(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": fmt.Sprintf("error when upload image: %s", err.Error()),
			"code":    400,
		})
	}

	var input dto.ProductCategoryRequestBody
	id := c.Params("id")
	idInt, _ := strconv.Atoi(id)

	input.Image = file
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
			"code":    400,
		})
	}

	admin := utils.CurrentUser(c)
	input.UpdatedBy = &admin.ID

	productCategory, err := h.s.UpdateProductCategory(idInt, input)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success update product category",
		"code":    200,
		"result":  productCategory,
	})
}

package admin

import (
	"strconv"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/utils"
)

type handler struct {
	s interfaces.AdminService
}

func NewAdminHandler(s interfaces.AdminService) interfaces.AdminHandler {
	return &handler{s}
}

// DeleteAdmin implements interfaces.AdminHandler.
func (h *handler) DeleteAdmin(c *fiber.Ctx) error {
	id := c.Params("id")
	ids, _ := strconv.Atoi(id)

	err := h.s.DeleteAdmin(ids)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success delete admin",
		"code":    200,
	})
}

// FindAllAdmin implements interfaces.AdminHandler.
func (h *handler) FindAllAdmin(c *fiber.Ctx) error {
	offset := c.Query("offset")
	limit := c.Query("limit")

	offsetInt, _ := strconv.Atoi(offset)
	limitInt, _ := strconv.Atoi(limit)

	admins, err := h.s.FindAllAdmin(offsetInt, limitInt)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success get all admins",
		"code":    200,
		"result":  admins,
	})
}

// FindOneAdminByID implements interfaces.AdminHandler.
func (h *handler) FindOneAdminByID(c *fiber.Ctx) error {
	id := c.Params("id")
	ids, _ := strconv.Atoi(id)

	admin, err := h.s.FindOneAdminByID(ids)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success get admin by id",
		"code":    200,
		"result":  admin,
	})
}

// InsertAdmin implements interfaces.AdminHandler.
func (h *handler) InsertAdmin(c *fiber.Ctx) error {
	var admin dto.AdminRequestBody
	if err := c.BodyParser(&admin); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to parse request body",
			"code":    400,
		})
	}

	user := utils.CurrentUser(c)
	admin.CreatedBy = &user.ID

	_, err := h.s.InsertAdmin(admin)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "success register admin",
		"code":    201,
	})
}

// UpdateAdmin implements interfaces.AdminHandler.
func (h *handler) UpdateAdmin(c *fiber.Ctx) error {
	id := c.Params("id")
	ids, _ := strconv.Atoi(id)

	var admin dto.AdminRequestBody
	if err := c.BodyParser(&admin); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to parse request body",
			"code":    400,
		})
	}

	user := utils.CurrentUser(c)
	admin.UpdatedBy = &user.ID

	_, err := h.s.UpdateAdmin(ids, admin)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success update admin",
		"code":    200,
	})
}

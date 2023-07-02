package user

import (
	"strconv"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/utils"
)

type handler struct {
	s interfaces.UserService
}

func NewUserHandler(s interfaces.UserService) interfaces.UserHandler {
	return &handler{s}
}

// InsertUser implements interfaces.UserHandler.
func (h *handler) InsertUser(c *fiber.Ctx) error {
	var input dto.UserRequestBody
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	admin := utils.CurrentUser(c)
	input.CreatedBy = &admin.ID

	user, err := h.s.InsertUser(&input)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"message": "success registered",
		"result":  user,
	})
}

// DeleteUser implements interfaces.UserHandler.
func (h *handler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	ids, _ := strconv.Atoi(id)

	err := h.s.DeleteUser(ids)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "success deleted",
	})
}

// FindAllUser implements interfaces.UserHandler.
func (h *handler) FindAllUser(c *fiber.Ctx) error {
	offset := c.Query("offset")
	limit := c.Query("limit")

	offsetInt, _ := strconv.Atoi(offset)
	limitInt, _ := strconv.Atoi(limit)

	users, err := h.s.FindAllUser(offsetInt, limitInt)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "success get all user",
		"result":  users,
	})
}

// FindUserById implements interfaces.UserHandler.
func (h *handler) FindUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	ids, _ := strconv.Atoi(id)

	user, err := h.s.FindUserById(ids)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "success get user by id",
		"result":  user,
	})
}

// UpdateUser implements interfaces.UserHandler.
func (h *handler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	ids, _ := strconv.Atoi(id)

	var input dto.UserUpdateRequestBody

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	admin := utils.CurrentUser(c)
	input.UpdatedBy = &admin.ID

	user, err := h.s.UpdateUser(ids, &input)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "success updated",
		"result":  user,
	})
}

package classroom

import (
	"strconv"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/utils"
)

type handler struct {
	s interfaces.ClassRoomService
}

func NewClassRoomHandler(s interfaces.ClassRoomService) interfaces.ClassRoomHandler {
	return &handler{s}
}

// FindAllClassRoomsByUserID implements interfaces.ClassRoomHandler.
func (h *handler) FindAllClassRoomsByUserID(c *fiber.Ctx) error {
	offset := c.Query("offset")
	limit := c.Query("limit")

	offsetInt, _ := strconv.Atoi(offset)
	limitInt, _ := strconv.Atoi(limit)

	user := utils.CurrentUser(c)

	res, err := h.s.FindAllClassRoomsByUserID(int(user.ID), offsetInt, limitInt)
	if err != nil {
		return c.Status(utils.GetStatusCode(err)).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success get all class room",
		"code":    fiber.StatusOK,
		"result":  res,
	})
}

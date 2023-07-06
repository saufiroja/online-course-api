package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/models/entity"
)

type ClassRoomRepository interface {
	FindAllClassRoomsByUserID(userID, offset, limit int) ([]entity.ClassRoom, error)
	FindOneClassRoomByUserIDAndProductID(userID, productID int) (*entity.ClassRoom, error)
	InsertClassRoom(input entity.ClassRoom) (entity.ClassRoom, error)
}

type ClassRoomService interface {
	FindAllClassRoomsByUserID(userID, offset, limit int) (*dto.ClassRoomListResponse, error)
	FindOneClassRoomByUserIDAndProductID(userID, productID int) (*dto.ClassRoomResponseBody, error)
	InsertClassRoom(input dto.ClassRoomRequestBody) (*entity.ClassRoom, error)
}

type ClassRoomHandler interface {
	FindAllClassRoomsByUserID(c *fiber.Ctx) error
}

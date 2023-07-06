package dto

import (
	"time"

	"github.com/saufiroja/online-course-api/models/entity"
	"gorm.io/gorm"
)

type ClassRoomRequestBody struct {
	UserID    int64 `json:"user_id"`
	ProductID int64 `json:"product_id"`
}

type ClassRoomResponseBody struct {
	ID        int64           `json:"id"`
	User      *entity.User    `json:"user"`
	Product   *entity.Product `json:"product"`
	CreatedBy *entity.User    `json:"created_by"`
	UpdatedBy *entity.User    `json:"updated_by"`
	CreatedAt *time.Time      `json:"created_at"`
	UpdatedAt *time.Time      `json:"updated_at"`
	DeleteAt  gorm.DeletedAt  `json:"deleted_at"`
}

func CreateClassRoomResponse(entity entity.ClassRoom) ClassRoomResponseBody {
	return ClassRoomResponseBody{
		ID:        entity.ID,
		User:      entity.User,
		Product:   entity.Product,
		CreatedBy: entity.CreatedBy,
		UpdatedBy: entity.UpdateBy,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeleteAt:  entity.DeletedAt,
	}
}

type ClassRoomListResponse []ClassRoomResponseBody

func CreateClassRoomListResponse(entity []entity.ClassRoom) ClassRoomListResponse {
	classRoomResp := ClassRoomListResponse{}

	for _, classRoom := range entity {
		classRoom.Product.VideoLink = classRoom.Product.Video

		classRoomResponse := CreateClassRoomResponse(classRoom)
		classRoomResp = append(classRoomResp, classRoomResponse)
	}

	return classRoomResp
}

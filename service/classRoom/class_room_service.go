package classroom

import (
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/models/entity"
	"github.com/saufiroja/online-course-api/utils"
)

type service struct {
	r interfaces.ClassRoomRepository
}

func NewClassRoomService(r interfaces.ClassRoomRepository) interfaces.ClassRoomService {
	return &service{r}
}

// FindAllClassRoomsByUserID implements interfaces.ClassRoomService.
func (s *service) FindAllClassRoomsByUserID(userID int, offset int, limit int) (*dto.ClassRoomListResponse, error) {
	classRoom, err := s.r.FindAllClassRoomsByUserID(userID, offset, limit)
	if err != nil {
		return nil, utils.HandlerError(404, "classroom not found")
	}

	classRes := dto.CreateClassRoomListResponse(classRoom)

	return &classRes, nil
}

// FindOneClassRoomByUserIDAndProductID implements interfaces.ClassRoomService.
func (s *service) FindOneClassRoomByUserIDAndProductID(userID int, productID int) (*dto.ClassRoomResponseBody, error) {
	classRoom, err := s.r.FindOneClassRoomByUserIDAndProductID(userID, productID)
	if err != nil {
		return nil, utils.HandlerError(404, "classroom not found")
	}

	classRes := dto.CreateClassRoomResponse(*classRoom)

	return &classRes, nil
}

// InsertClassRoom implements interfaces.ClassRoomService.
func (s *service) InsertClassRoom(input dto.ClassRoomRequestBody) (*entity.ClassRoom, error) {
	classRoom, err := s.r.FindOneClassRoomByUserIDAndProductID(int(input.UserID), int(input.ProductID))
	if err != nil {
		return nil, utils.HandlerError(404, "classroom not found")
	}

	if classRoom != nil {
		return nil, utils.HandlerError(400, "classroom already exists")
	}

	res := entity.ClassRoom{
		UserID:      input.UserID,
		ProductID:   &input.ProductID,
		CreatedByID: &input.UserID,
	}

	result, err := s.r.InsertClassRoom(res)
	if err != nil {
		return nil, utils.HandlerError(500, "failed to insert classroom")
	}

	return &result, nil
}

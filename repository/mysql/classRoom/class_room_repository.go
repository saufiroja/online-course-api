package classroom

import (
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/entity"
	"github.com/saufiroja/online-course-api/utils"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewClassRoomRepository(db *gorm.DB) interfaces.ClassRoomRepository {
	return &repository{db}
}

// FindAllClassRoomsByUserID implements interfaces.ClassRoomRepository.
func (r *repository) FindAllClassRoomsByUserID(userID int, offset int, limit int) ([]entity.ClassRoom, error) {
	var classRooms []entity.ClassRoom

	err := r.db.Scopes(utils.Paginate(offset, limit)).
		Preload("Product.ProductCategory").
		Where("user_id = ?", userID).
		Find(&classRooms).Error
	if err != nil {
		return classRooms, err
	}

	return classRooms, nil
}

// FindOneClassRoomByUserIDAndProductID implements interfaces.ClassRoomRepository.
func (r *repository) FindOneClassRoomByUserIDAndProductID(userID int, productID int) (*entity.ClassRoom, error) {
	var classRoom entity.ClassRoom

	err := r.db.Preload("Product.ProductCategory").
		Where("user_id = ?", userID).
		Where("product_id = ?", productID).
		First(&classRoom).Error

	if err != nil {
		return nil, err
	}

	return &classRoom, nil
}

// InsertClassRoom implements interfaces.ClassRoomRepository.
func (r *repository) InsertClassRoom(input entity.ClassRoom) (entity.ClassRoom, error) {
	err := r.db.Create(&input).Error
	if err != nil {
		return input, err
	}

	return input, nil
}

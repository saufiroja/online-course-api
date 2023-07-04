package orderdetail

import (
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/entity"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewOrderDetailRepository(db *gorm.DB) interfaces.OrderDetailRepository {
	return &repository{
		db: db,
	}
}

// InsertOrderDetail implements interfaces.OrderDetailRepository.
func (r *repository) InsertOrderDetail(input entity.OrderDetail) error {
	err := r.db.Create(&input).Error
	if err != nil {
		return err
	}

	return nil
}

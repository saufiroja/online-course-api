package order

import (
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/entity"
	"github.com/saufiroja/online-course-api/utils"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) interfaces.OrderRepository {
	return &repository{
		db: db,
	}
}

// CountOrder implements interfaces.OrderRepository.
func (r *repository) CountOrder() (int64, error) {
	var count int64
	var orders entity.Order

	err := r.db.Model(&orders).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// FindAllOrdersByUserId implements interfaces.OrderRepository.
func (r *repository) FindAllOrdersByUserId(userId int, offset int, limit int) ([]entity.Order, error) {
	var orders []entity.Order
	err := r.db.Scopes(utils.Paginate(offset, limit)).
		Preload("OrderDetails.Product").
		Where("user_id = ?", userId).
		Find(&orders).Error
	if err != nil {
		return orders, err
	}

	return orders, nil
}

// FindOrderByExternalId implements interfaces.OrderRepository.
func (r *repository) FindOrderByExternalId(externalId string) (entity.Order, error) {
	var order entity.Order
	err := r.db.Preload("OrderDetails.Product").
		Where("external_id = ?", externalId).
		First(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

// FindOrderById implements interfaces.OrderRepository.
func (r *repository) FindOrderById(id int) (entity.Order, error) {
	var order entity.Order
	err := r.db.Preload("OrderDetails.Product").
		First(&order, id).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

// InsertOrder implements interfaces.OrderRepository.
func (r *repository) InsertOrder(input entity.Order) (entity.Order, error) {
	err := r.db.Create(&input).Error
	if err != nil {
		return input, err
	}

	return input, nil
}

// UpdateOrder implements interfaces.OrderRepository.
func (r *repository) UpdateOrder(input entity.Order) (entity.Order, error) {
	err := r.db.Save(&input).Error
	if err != nil {
		return input, err
	}

	return input, nil
}

package cart

import (
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/entity"
	"github.com/saufiroja/online-course-api/utils"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) interfaces.CartRepository {
	return &repository{
		db: db,
	}
}

// DeleteCart implements interfaces.CartRepository.
func (r *repository) DeleteCart(input entity.Cart) error {
	err := r.db.Delete(&input).Error
	if err != nil {
		return err
	}

	return nil
}

// DeleteCartByUserId implements interfaces.CartRepository.
func (r *repository) DeleteCartByUserId(userId int) error {
	var cart entity.Cart
	err := r.db.Where("user_id = ?", userId).Delete(&cart).Error
	if err != nil {
		return err
	}

	return nil
}

// FindAllCartByUserId implements interfaces.CartRepository.
func (r *repository) FindAllCartByUserId(userId int, offset int, limit int) ([]entity.Cart, error) {
	var carts []entity.Cart
	err := r.db.Scopes(utils.Paginate(offset, limit)).
		Preload("User").Preload("Product").
		Where("user_id = ?", userId).
		Find(&carts).Error

	if err != nil {
		return nil, err
	}

	return carts, nil
}

// FindCartById implements interfaces.CartRepository.
func (r *repository) FindCartById(id int) (*entity.Cart, error) {
	var cart entity.Cart
	err := r.db.Preload("User").Preload("Product").Find(&cart, id).Error
	if err != nil {
		return nil, err
	}

	return &cart, nil
}

// InsertCart implements interfaces.CartRepository.
func (r *repository) InsertCart(input entity.Cart) (*entity.Cart, error) {
	err := r.db.Create(&input).Error
	if err != nil {
		return nil, err
	}

	return &input, nil
}

// UpdateCart implements interfaces.CartRepository.
func (r *repository) UpdateCart(input entity.Cart) (*entity.Cart, error) {
	err := r.db.Save(&input).Error
	if err != nil {
		return nil, err
	}

	return &input, nil
}

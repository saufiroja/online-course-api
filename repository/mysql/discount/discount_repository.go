package discount

import (
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/entity"
	"github.com/saufiroja/online-course-api/utils"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewDiscountRepository(db *gorm.DB) interfaces.DiscountRepository {
	return &repository{db}
}

// DeleteDiscountById implements interfaces.DiscountRepository.
func (r *repository) DeleteDiscountById(input entity.Discount) error {
	err := r.db.Delete(&input).Error
	if err != nil {
		return err
	}

	return nil
}

// FindAllDiscount implements interfaces.DiscountRepository.
func (r *repository) FindAllDiscount(offset int, limit int) ([]entity.Discount, error) {
	var discounts []entity.Discount
	err := r.db.Scopes(utils.Paginate(offset, limit)).Find(&discounts).Error
	if err != nil {
		return nil, err
	}

	return discounts, nil
}

// FindDiscountByCode implements interfaces.DiscountRepository.
func (r *repository) FindDiscountByCode(code string) (*entity.Discount, error) {
	var discount entity.Discount
	err := r.db.Where("code = ?", code).First(&discount).Error
	if err != nil {
		return nil, err
	}

	return &discount, nil
}

// FindDiscountById implements interfaces.DiscountRepository.
func (r *repository) FindDiscountById(id int) (*entity.Discount, error) {
	var discount entity.Discount
	err := r.db.First(&discount, id).Error
	if err != nil {
		return nil, err
	}

	return &discount, nil
}

// InsertDiscount implements interfaces.DiscountRepository.
func (r *repository) InsertDiscount(input entity.Discount) (*entity.Discount, error) {
	err := r.db.Create(&input).Error
	if err != nil {
		return nil, err
	}

	return &input, nil
}

// UpdateDiscountById implements interfaces.DiscountRepository.
func (r *repository) UpdateDiscountById(input entity.Discount) (*entity.Discount, error) {
	err := r.db.Save(&input).Error
	if err != nil {
		return nil, err
	}

	return &input, nil
}

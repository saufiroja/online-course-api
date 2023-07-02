package product

import (
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/entity"
	"github.com/saufiroja/online-course-api/utils"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) interfaces.ProductRepository {
	return &repository{db}
}

// CountProduct implements interfaces.ProductRepository.
func (r *repository) CountProduct() (int64, error) {
	var count int64
	if err := r.db.Model(&entity.Product{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// DeleteProductByID implements interfaces.ProductRepository.
func (r *repository) DeleteProductByID(input entity.Product) error {
	err := r.db.Delete(&input).Error
	if err != nil {
		return err
	}

	return nil
}

// FindAllProduct implements interfaces.ProductRepository.
func (r *repository) FindAllProduct(offset int, limit int) ([]entity.Product, error) {
	var products []entity.Product
	err := r.db.Scopes(utils.Paginate(offset, limit)).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

// FindProductByID implements interfaces.ProductRepository.
func (r *repository) FindProductByID(id int) (*entity.Product, error) {
	var product entity.Product
	err := r.db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// InsertProduct implements interfaces.ProductRepository.
func (r *repository) InsertProduct(input entity.Product) (*entity.Product, error) {
	err := r.db.Create(&input).Error
	if err != nil {
		return nil, err
	}

	return &input, nil
}

// UpdateProductByID implements interfaces.ProductRepository.
func (r *repository) UpdateProductByID(input entity.Product) (*entity.Product, error) {
	err := r.db.Save(&input).Error
	if err != nil {
		return nil, err
	}

	return &input, nil
}

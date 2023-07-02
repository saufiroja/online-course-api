package productcategory

import (
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/entity"
	"github.com/saufiroja/online-course-api/utils"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewProductCategoryRepository(db *gorm.DB) interfaces.ProductCategoryRepository {
	return &repository{db}
}

// DeleteProductCategory implements interfaces.ProductCategoryRepository.
func (r *repository) DeleteProductCategory(input entity.ProductCategory) error {
	err := r.db.Delete(&input).Error
	if err != nil {
		return err
	}

	return nil
}

// FindAllProductCategory implements interfaces.ProductCategoryRepository.
func (r *repository) FindAllProductCategory(offset int, limit int) ([]entity.ProductCategory, error) {
	var productCategories []entity.ProductCategory
	err := r.db.Scopes(utils.Paginate(offset, limit)).Find(&productCategories).Error
	if err != nil {
		return nil, err
	}

	return productCategories, nil
}

// FindProductCategoryByID implements interfaces.ProductCategoryRepository.
func (r *repository) FindProductCategoryByID(id int) (entity.ProductCategory, error) {
	var productCategory entity.ProductCategory
	err := r.db.First(&productCategory, id).Error
	if err != nil {
		return productCategory, err
	}

	return productCategory, nil
}

// InsertProductCategory implements interfaces.ProductCategoryRepository.
func (r *repository) InsertProductCategory(input entity.ProductCategory) (entity.ProductCategory, error) {
	err := r.db.Create(&input).Error
	if err != nil {
		return input, err
	}

	return input, nil
}

// UpdateProductCategory implements interfaces.ProductCategoryRepository.
func (r *repository) UpdateProductCategory(input entity.ProductCategory) (entity.ProductCategory, error) {
	err := r.db.Save(&input).Error
	if err != nil {
		return input, err
	}

	return input, nil
}

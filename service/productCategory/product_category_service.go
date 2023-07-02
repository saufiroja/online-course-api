package productcategory

import (
	"github.com/saufiroja/online-course-api/config"
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/models/entity"
	"github.com/saufiroja/online-course-api/utils"
)

type service struct {
	r    interfaces.ProductCategoryRepository
	conf *config.AppConfig
}

func NewProductCategoryService(r interfaces.ProductCategoryRepository, conf *config.AppConfig) interfaces.ProductCategoryService {
	return &service{
		r,
		conf,
	}
}

// DeleteProductCategory implements interfaces.ProductCategoryService.
func (s *service) DeleteProductCategory(id int) error {
	category, err := s.r.FindProductCategoryByID(id)
	if err != nil {
		return utils.HandlerError(404, "Product category not found")
	}

	err = s.r.DeleteProductCategory(category)
	if err != nil {
		return utils.HandlerError(500, err.Error())
	}

	return nil
}

// FindAllProductCategory implements interfaces.ProductCategoryService.
func (s *service) FindAllProductCategory(offset int, limit int) ([]entity.ProductCategory, error) {
	res, err := s.r.FindAllProductCategory(offset, limit)
	if err != nil {
		return nil, utils.HandlerError(500, err.Error())
	}

	return res, nil
}

// FindProductCategoryByID implements interfaces.ProductCategoryService.
func (s *service) FindProductCategoryByID(id int) (entity.ProductCategory, error) {
	res, err := s.r.FindProductCategoryByID(id)
	if err != nil {
		return res, utils.HandlerError(404, "Product category not found")
	}

	return res, nil
}

// InsertProductCategory implements interfaces.ProductCategoryService.
func (s *service) InsertProductCategory(input dto.ProductCategoryRequestBody) (entity.ProductCategory, error) {
	data := entity.ProductCategory{
		Name:        input.Name,
		CreatedByID: input.CreatedBy,
	}
	if input.Image != nil {
		image, err := utils.UploadImage(*input.Image, s.conf)
		if err != nil {
			return data, utils.HandlerError(500, err.Error())
		}

		if image != nil {
			data.Image = image
		}
	}

	res, err := s.r.InsertProductCategory(data)
	if err != nil {
		return res, utils.HandlerError(500, err.Error())
	}

	return res, nil
}

// UpdateProductCategory implements interfaces.ProductCategoryService.
func (s *service) UpdateProductCategory(id int, input dto.ProductCategoryRequestBody) (entity.ProductCategory, error) {
	category, err := s.r.FindProductCategoryByID(id)
	if err != nil {
		return category, utils.HandlerError(404, "Product category not found")
	}

	category.Name = input.Name
	category.UpdatedByID = input.UpdatedBy

	if input.Image != nil {
		image, err := utils.UploadImage(*input.Image, s.conf)
		if err != nil {
			return category, utils.HandlerError(500, err.Error())
		}

		if category.Image != nil {
			_, err := utils.DeleteImage(*category.Image, s.conf)
			if err != nil {
				return category, utils.HandlerError(500, err.Error())
			}
		}

		if image != nil {
			category.Image = image
		}
	}

	res, err := s.r.UpdateProductCategory(category)
	if err != nil {
		return res, utils.HandlerError(500, err.Error())
	}

	return res, nil
}

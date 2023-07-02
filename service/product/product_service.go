package product

import (
	"github.com/saufiroja/online-course-api/config"
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/models/entity"
	"github.com/saufiroja/online-course-api/utils"
)

type service struct {
	r    interfaces.ProductRepository
	conf *config.AppConfig
}

func NewProductService(r interfaces.ProductRepository, conf *config.AppConfig) interfaces.ProductServices {
	return &service{
		r,
		conf,
	}
}

// CountProduct implements interfaces.ProductServices.
func (s *service) CountProduct() (int64, error) {
	count, err := s.r.CountProduct()
	if err != nil {
		return 0, utils.HandlerError(500, err.Error())
	}

	return count, nil
}

// DeleteProductByID implements interfaces.ProductServices.
func (s *service) DeleteProductByID(id int) error {
	product, err := s.r.FindProductByID(id)
	if err != nil {
		return utils.HandlerError(404, "Product not found")
	}

	err = s.r.DeleteProductByID(*product)
	if err != nil {
		return utils.HandlerError(500, err.Error())
	}

	return nil
}

// FindAllProduct implements interfaces.ProductServices.
func (s *service) FindAllProduct(offset int, limit int) ([]entity.Product, error) {
	product, err := s.r.FindAllProduct(offset, limit)
	if err != nil {
		return nil, utils.HandlerError(500, err.Error())
	}

	return product, nil
}

// FindProductByID implements interfaces.ProductServices.
func (s *service) FindProductByID(id int) (*entity.Product, error) {
	product, err := s.r.FindProductByID(id)
	if err != nil {
		return nil, utils.HandlerError(404, "Product not found")
	}

	return product, nil
}

// InsertProduct implements interfaces.ProductServices.
func (s *service) InsertProduct(input dto.ProductRequestBody) (*entity.Product, error) {
	product := entity.Product{
		ProductCategoryID: &input.ProductCategoryID,
		Title:             input.Title,
		Description:       input.Description,
		IsHighlighted:     input.IsHighlighted,
		Price:             int64(input.Price),
		CreatedByID:       input.CreatedBy,
	}

	if input.Image != nil {
		image, err := utils.UploadImage(*input.Image, s.conf)
		if err != nil {
			return nil, utils.HandlerError(500, err.Error())
		}

		if image != nil {
			product.Image = image
		}
	}

	if input.Video != nil {
		video, err := utils.UploadImage(*input.Video, s.conf)
		if err != nil {
			return nil, utils.HandlerError(500, err.Error())
		}

		if video != nil {
			product.Video = video
		}
	}

	res, err := s.r.InsertProduct(product)
	if err != nil {
		return nil, utils.HandlerError(500, err.Error())
	}

	return res, nil
}

// UpdateProductByID implements interfaces.ProductServices.
func (s *service) UpdateProductByID(id int, input dto.ProductRequestBody) (*entity.Product, error) {
	product, err := s.r.FindProductByID(id)
	if err != nil {
		return nil, utils.HandlerError(404, "Product not found")
	}

	product.ProductCategoryID = &input.ProductCategoryID
	product.Title = input.Title
	product.Description = input.Description
	product.IsHighlighted = input.IsHighlighted
	product.Price = int64(input.Price)
	product.UpdatedByID = input.UpdatedBy

	if input.Image != nil {
		image, err := utils.UploadImage(*input.Image, s.conf)
		if err != nil {
			return nil, utils.HandlerError(500, err.Error())
		}

		if product.Image != nil {
			_, err = utils.DeleteImage(*product.Image, s.conf)
			if err != nil {
				return nil, utils.HandlerError(500, err.Error())
			}
		}

		if image != nil {
			product.Image = image
		}
	}

	if input.Video != nil {
		video, err := utils.UploadImage(*input.Video, s.conf)
		if err != nil {
			return nil, utils.HandlerError(500, err.Error())
		}

		if product.Video != nil {
			_, err = utils.DeleteImage(*product.Video, s.conf)
			if err != nil {
				return nil, utils.HandlerError(500, err.Error())
			}
		}

		if video != nil {
			product.Video = video
		}
	}

	res, err := s.r.UpdateProductByID(*product)
	if err != nil {
		return nil, utils.HandlerError(500, err.Error())
	}

	return res, nil
}

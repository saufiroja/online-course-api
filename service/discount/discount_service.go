package discount

import (
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/models/entity"
	"github.com/saufiroja/online-course-api/utils"
)

type service struct {
	r interfaces.DiscountRepository
}

func NewDiscountService(r interfaces.DiscountRepository) interfaces.DiscountService {
	return &service{r}
}

// DeleteDiscountById implements interfaces.DiscountService.
func (s *service) DeleteDiscountById(id int) error {
	discount, err := s.r.FindDiscountById(id)
	if err != nil {
		return utils.HandlerError(404, "discount not found")
	}

	err = s.r.DeleteDiscountById(*discount)
	if err != nil {
		return utils.HandlerError(500, "failed to delete discount")
	}

	return nil
}

// FindAllDiscount implements interfaces.DiscountService.
func (s *service) FindAllDiscount(offset int, limit int) ([]entity.Discount, error) {
	discount, err := s.r.FindAllDiscount(offset, limit)
	if err != nil {
		return nil, utils.HandlerError(500, "failed to get discounts")
	}

	return discount, nil
}

// FindDiscountByCode implements interfaces.DiscountService.
func (s *service) FindDiscountByCode(code string) (*entity.Discount, error) {
	discount, err := s.r.FindDiscountByCode(code)
	if err != nil {
		return nil, utils.HandlerError(404, "discount not found")
	}

	return discount, nil
}

// FindDiscountById implements interfaces.DiscountService.
func (s *service) FindDiscountById(id int) (*entity.Discount, error) {
	discount, err := s.r.FindDiscountById(id)
	if err != nil {
		return nil, utils.HandlerError(404, "discount not found")
	}

	return discount, nil
}

// InsertDiscount implements interfaces.DiscountService.
func (s *service) InsertDiscount(input dto.DiscountRequestBody) (*entity.Discount, error) {
	discount := entity.Discount{
		Code:              input.Code,
		Name:              input.Name,
		Quantity:          input.Quantity,
		RemainingQuantity: input.Quantity,
		Type:              input.Type,
		Value:             input.Value,
		CreatedByID:       input.CreatedBy,
	}

	if input.StartDate != nil {
		discount.StartDate = input.StartDate
	}

	if input.EndDate != nil {
		discount.EndDate = input.EndDate
	}

	res, err := s.r.InsertDiscount(discount)
	if err != nil {
		return nil, utils.HandlerError(500, "failed to insert discount")
	}

	return res, nil
}

// UpdateDiscounById implements interfaces.DiscountService.
func (s *service) UpdateDiscountById(id int, input dto.DiscountRequestBody) (*entity.Discount, error) {
	discount, err := s.r.FindDiscountById(id)
	if err != nil {
		return nil, utils.HandlerError(404, "discount not found")
	}

	discount.Name = input.Name
	discount.Code = input.Code
	discount.Quantity = input.Quantity
	discount.RemainingQuantity = input.Quantity
	discount.Type = input.Type
	discount.Value = input.Value

	if input.StartDate != nil {
		discount.StartDate = input.StartDate
	}

	if input.EndDate != nil {
		discount.EndDate = input.EndDate
	}

	res, err := s.r.UpdateDiscountById(*discount)
	if err != nil {
		return nil, utils.HandlerError(500, "failed to update discount")
	}

	return res, nil
}

// UpdateRemainingDiscount implements interfaces.DiscountService.
func (s *service) UpdateRemainingDiscount(id int, input dto.DiscountRemainingQuantityRequestBody) (*entity.Discount, error) {
	discount, err := s.r.FindDiscountById(id)
	if err != nil {
		return nil, utils.HandlerError(404, "discount not found")
	}

	if input.Operator == "+" {
		discount.RemainingQuantity += input.Quantity
	} else if input.Operator == "-" {
		discount.RemainingQuantity -= input.Quantity
	}

	updateDiscount, err := s.r.UpdateDiscountById(*discount)
	if err != nil {
		return nil, utils.HandlerError(500, "failed to update discount")
	}

	return updateDiscount, nil
}

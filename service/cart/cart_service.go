package cart

import (
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/models/entity"
	"github.com/saufiroja/online-course-api/utils"
)

type service struct {
	r interfaces.CartRepository
}

func NewCartService(r interfaces.CartRepository) interfaces.CartService {
	return &service{
		r: r,
	}
}

// DeleteCart implements interfaces.CartService.
func (s *service) DeleteCart(id, userId int) error {
	cart, err := s.r.FindCartById(id)
	if err != nil {
		return utils.HandlerError(404, "Cart not found")
	}

	user := int64(userId)

	if *cart.UserID != user {
		return utils.HandlerError(400, "You are not authorized to delete this cart")
	}

	err = s.r.DeleteCart(*cart)
	if err != nil {
		return utils.HandlerError(500, err.Error())
	}

	return nil
}

// DeleteCartByUserId implements interfaces.CartService.
func (s *service) DeleteCartByUserId(userId int, input dto.CartRequestUpdateBody) error {
	err := s.r.DeleteCartByUserId(userId)
	if err != nil {
		return utils.HandlerError(404, "Cart not found")
	}

	return nil
}

// FindAllCartByUserId implements interfaces.CartService.
func (s *service) FindAllCartByUserId(userId int, offset int, limit int) ([]entity.Cart, error) {
	cart, err := s.r.FindAllCartByUserId(userId, offset, limit)
	if err != nil {
		return nil, utils.HandlerError(404, "Cart not found")
	}

	return cart, nil
}

// FindCartById implements interfaces.CartService.
func (s *service) FindCartById(id int) (*entity.Cart, error) {
	res, err := s.r.FindCartById(id)
	if err != nil {
		return nil, utils.HandlerError(404, "Cart not found")
	}

	return res, nil
}

// InsertCart implements interfaces.CartService.
func (s *service) InsertCart(input dto.CartRequestBody) (*entity.Cart, error) {
	cart := entity.Cart{
		UserID:      &input.UserID,
		ProductID:   &input.ProductID,
		Quantity:    1,
		IsChecked:   true,
		CreatedByID: &input.CreatedBy,
	}

	res, err := s.r.InsertCart(cart)
	if err != nil {
		return nil, utils.HandlerError(500, err.Error())
	}

	return res, nil
}

// UpdateCart implements interfaces.CartService.
func (s *service) UpdateCart(id int, input dto.CartRequestUpdateBody) (*entity.Cart, error) {
	cart, err := s.r.FindCartById(id)
	if err != nil {
		return nil, utils.HandlerError(404, "Cart not found")
	}

	if *cart.UserID != *input.UserID {
		return nil, utils.HandlerError(400, "You are not authorized to update this cart")
	}

	cart.IsChecked = input.IsChecked
	cart.UpdatedByID = input.UserID

	res, err := s.r.UpdateCart(*cart)
	if err != nil {
		return nil, utils.HandlerError(500, err.Error())
	}

	return res, nil
}

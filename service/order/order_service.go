package order

import (
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/models/entity"
	"github.com/saufiroja/online-course-api/utils"
)

type service struct {
	r              interfaces.OrderRepository
	cartSvc        interfaces.CartService
	discountSvc    interfaces.DiscountService
	productSvc     interfaces.ProductServices
	orderDetailSvc interfaces.OrderDetailService
	payment        interfaces.PaymentService
}

func NewOrderService(
	r interfaces.OrderRepository,
	cartSvc interfaces.CartService,
	discountSvc interfaces.DiscountService,
	productSvc interfaces.ProductServices,
	orderDetailSvc interfaces.OrderDetailService,
	payment interfaces.PaymentService) interfaces.OrderService {
	return &service{
		r:              r,
		cartSvc:        cartSvc,
		discountSvc:    discountSvc,
		productSvc:     productSvc,
		orderDetailSvc: orderDetailSvc,
		payment:        payment,
	}
}

// CountOrder implements interfaces.OrderService.
func (s *service) CountOrder() (int64, error) {
	count, err := s.r.CountOrder()
	if err != nil {
		return 0, utils.HandlerError(500, err.Error())
	}

	return count, nil
}

// FindAllOrdersByUserId implements interfaces.OrderService.
func (s *service) FindAllOrdersByUserId(userId int, offset int, limit int) ([]entity.Order, error) {
	order, err := s.r.FindAllOrdersByUserId(userId, offset, limit)
	if err != nil {
		return order, utils.HandlerError(500, err.Error())
	}

	return order, nil
}

// FindOrderByExternalId implements interfaces.OrderService.
func (s *service) FindOrderByExternalId(externalId string) (entity.Order, error) {
	order, err := s.r.FindOrderByExternalId(externalId)
	if err != nil {
		return order, utils.HandlerError(404, "order not found")
	}

	return order, nil
}

// FindOrderById implements interfaces.OrderService.
func (s *service) FindOrderById(id int) (entity.Order, error) {
	order, err := s.r.FindOrderById(id)
	if err != nil {
		return order, utils.HandlerError(404, "order not found")
	}

	return order, nil
}

// InsertOrder implements interfaces.OrderService.
func (s *service) InsertOrder(input dto.OrderRequestBody) (entity.Order, error) {
	price := 0
	totalPrice := 0
	description := ""

	var products []entity.Product
	var dataDiscount *entity.Discount

	order := entity.Order{
		UserID: &input.UserID,
		Status: "pending",
	}

	carts, err := s.cartSvc.FindAllCartByUserId(int(*order.UserID), 1, 100)
	if err != nil {
		return order, utils.HandlerError(500, err.Error())
	}

	if len(carts) == 0 {
		if input.ProductID == nil {
			return order, utils.HandlerError(400, "cart is empty")
		}
	}

	if input.DiscountCode != nil {
		discount, err := s.discountSvc.FindDiscountByCode(*input.DiscountCode)
		if err != nil {
			return order, utils.HandlerError(404, "discount not found")
		}

		if discount.RemainingQuantity == 0 {
			return order, utils.HandlerError(400, "discount is out of stock")
		}

		if discount.StartDate != nil && discount.EndDate != nil {
			if discount.StartDate.After(time.Now()) || discount.EndDate.Before(time.Now()) {
				return order, utils.HandlerError(400, "discount start date is after end date")
			}
		} else if discount.StartDate != nil {
			if discount.StartDate.After(time.Now()) {
				return order, utils.HandlerError(400, "discount start date is after end date")
			}
		} else if discount.EndDate != nil {
			if discount.EndDate.Before(time.Now()) {
				return order, utils.HandlerError(400, "discount start date is after end date")
			}
		}

		dataDiscount = discount
	}

	if len(carts) > 0 {
		for _, cart := range carts {
			product, err := s.productSvc.FindProductByID(int(*cart.ProductID))
			if err != nil {
				return order, utils.HandlerError(404, "product not found")
			}

			products = append(products, *product)
		}
	} else if input.ProductID != nil {
		product, err := s.productSvc.FindProductByID(int(*input.ProductID))
		if err != nil {
			return order, utils.HandlerError(404, "product not found")
		}

		products = append(products, *product)
	}

	for i, product := range products {
		price += int(product.Price)

		index := strconv.Itoa(i + 1)

		description += fmt.Sprintf("%s. Product : %s\n", index, product.Title)
	}

	totalPrice = price

	if dataDiscount != nil {
		if dataDiscount.Type == "rebate" {
			totalPrice = price - int(dataDiscount.Value)
		} else if dataDiscount.Type == "percentage" {
			totalPrice = price - (int(dataDiscount.Value) * price / 100)
		}
	}

	externarId := uuid.New().String()

	order.Price = int64(price)
	order.TotalPrice = int64(totalPrice)
	order.CreatedByID = &input.UserID
	order.ExternalID = externarId

	data, err := s.r.InsertOrder(order)
	if err != nil {
		return data, utils.HandlerError(500, err.Error())
	}

	for _, product := range products {
		orderDetail := entity.OrderDetail{
			ProductID:   &product.ID,
			Price:       product.Price,
			CreatedByID: &input.UserID,
			OrderID:     data.ID,
		}

		err := s.orderDetailSvc.InsertOrderDetail(orderDetail)
		if err != nil {
			return data, utils.HandlerError(500, err.Error())
		}
	}

	// payment
	payment := dto.PaymentRequestBody{
		ExternalID:  externarId,
		Amount:      int64(totalPrice),
		PayerEmail:  input.Email,
		Description: description,
	}

	res, err := s.payment.InsertPayment(payment)
	if err != nil {
		return data, utils.HandlerError(500, err.Error())
	}

	data.CheckoutLink = res.InvoiceURL

	if input.DiscountCode != nil {
		dto := dto.DiscountRemainingQuantityRequestBody{
			Quantity: 1,
			Operator: "-",
		}
		_, err := s.discountSvc.UpdateRemainingDiscount(int(dataDiscount.ID), dto)
		if err != nil {
			return data, utils.HandlerError(500, err.Error())
		}
	}

	err = s.cartSvc.DeleteCartByUserId(int(*order.UserID))
	if err != nil {
		return data, utils.HandlerError(500, err.Error())
	}

	return data, nil
}

// UpdateOrder implements interfaces.OrderService.
func (s *service) UpdateOrder(id int, input dto.OrderRequestBody) (entity.Order, error) {
	order, err := s.r.FindOrderById(id)
	if err != nil {
		return order, utils.HandlerError(404, "order not found")
	}

	order.Status = input.Status

	res, err := s.r.UpdateOrder(order)
	if err != nil {
		return res, utils.HandlerError(500, err.Error())
	}

	return res, nil
}

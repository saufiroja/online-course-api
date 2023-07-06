package webhook

import (
	"strings"

	"github.com/saufiroja/online-course-api/config"
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/utils"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

type service struct {
	orderSvc     interfaces.OrderService
	classRoomSvc interfaces.ClassRoomService
	conf         *config.AppConfig
}

func NewWebhookService(orderSvc interfaces.OrderService, classRoomSvc interfaces.ClassRoomService, conf *config.AppConfig) interfaces.WebhookService {
	return &service{
		orderSvc:     orderSvc,
		classRoomSvc: classRoomSvc,
		conf:         conf,
	}
}

// UpdatePayment implements interfaces.WebhookService.
func (s *service) UpdatePayment(id string) error {
	xendit.Opt.SecretKey = s.conf.Xendit.SecretKey
	param := invoice.GetParams{
		ID: id,
	}

	resp, err := invoice.Get(&param)
	if err != nil {
		return utils.HandlerError(400, err.Error())
	}

	if resp == nil {
		return utils.HandlerError(400, "invoice not found")
	}

	order, errs := s.orderSvc.FindOrderByExternalId(resp.ExternalID)
	if errs != nil {
		return utils.HandlerError(400, errs.Error())
	}

	if order == nil {
		return utils.HandlerError(400, "order not found")
	}

	if order.Status == "settled" {
		return utils.HandlerError(400, "order already paid")
	}

	if order.Status != "paid" {
		if resp.Status == "PAID" || resp.Status == "SETTLED" {
			for _, v := range order.OrderDetails {
				classRoom := dto.ClassRoomRequestBody{
					UserID:    *order.UserID,
					ProductID: *v.ProductID,
				}

				_, err := s.classRoomSvc.InsertClassRoom(classRoom)
				if err != nil {
					return utils.HandlerError(400, err.Error())
				}
			}
		}
	}

	req := dto.OrderRequestBody{
		Status: strings.ToLower(resp.Status),
	}

	_, errr := s.orderSvc.UpdateOrder(int(order.ID), req)
	if errr != nil {
		return utils.HandlerError(400, err.Error())
	}

	return nil
}

package interfaces

import (
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/xendit/xendit-go"
)

type PaymentService interface {
	InsertPayment(input dto.PaymentRequestBody) (*xendit.Invoice, error)
}

package payment

import (
	"github.com/saufiroja/online-course-api/config"
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/utils"
	xendit "github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

type service struct {
	conf *config.AppConfig
}

func NewPaymentService(conf *config.AppConfig) interfaces.PaymentService {
	return &service{
		conf: conf,
	}
}

// CreatePayment implements interfaces.PaymentService.
func (s *service) CreatePayment(input dto.PaymentRequestBody) (*xendit.Invoice, error) {
	xendit.Opt.SecretKey = s.conf.Xendit.SecretKey
	url := s.conf.Xendit.SuccessUrl

	data := invoice.CreateParams{
		ExternalID:  input.ExternalID,
		Amount:      float64(input.Amount),
		Description: input.Description,
		PayerEmail:  input.PayerEmail,
		Customer: xendit.InvoiceCustomer{
			Email: input.PayerEmail,
		},
		CustomerNotificationPreference: xendit.InvoiceCustomerNotificationPreference{
			InvoiceCreated:  []string{"EMAIL"},
			InvoicePaid:     []string{"EMAIL"},
			InvoiceReminder: []string{"EMAIL"},
			InvoiceExpired:  []string{"EMAIL"},
		},
		InvoiceDuration:    86400,
		SuccessRedirectURL: url,
		FailureRedirectURL: url,
	}

	resp, err := invoice.Create(&data)
	if err != nil {
		return nil, utils.HandlerError(500, err.Error())
	}

	return resp, nil
}

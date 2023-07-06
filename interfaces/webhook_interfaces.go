package interfaces

type WebhookService interface {
	UpdatePayment(id string) error
}

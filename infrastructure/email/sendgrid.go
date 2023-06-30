package email

import (
	"log"

	"github.com/saufiroja/online-course-api/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func NewSendEmail(toEmail, result, subject string, conf *config.AppConfig) {
	name := conf.SendGrid.MailSenderName
	key := conf.SendGrid.MailSenderKey

	from := mail.NewEmail(name, name)
	to := mail.NewEmail(toEmail, toEmail)

	message := mail.NewSingleEmail(from, subject, to, "", result)
	client := sendgrid.NewSendClient(key)
	res, err := client.Send(message)
	if err != nil {
		panic(err)
	}

	if res.StatusCode != 202 {
		panic(err)
	}

	log.Println("success send email!")
}

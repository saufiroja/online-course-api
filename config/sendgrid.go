package config

import "os"

func initSendGrid(conf *AppConfig) {
	mailName := os.Getenv("MAIL_SENDER_NAME")
	mailKey := os.Getenv("MAIL_SENDER_KEY")

	conf.SendGrid.MailSenderName = mailName
	conf.SendGrid.MailSenderKey = mailKey
}

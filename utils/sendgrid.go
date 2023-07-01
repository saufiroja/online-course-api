package utils

import (
	"bytes"

	"github.com/saufiroja/online-course-api/config"
	"github.com/saufiroja/online-course-api/infrastructure/email"
	"github.com/saufiroja/online-course-api/models/dto"
)

func SendVerification(toEmail string, data dto.EmailVerification, conf *config.AppConfig) {
	t, err := templateRender.templates.ParseFS(views, "template/verification_email.html")
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		panic(err)
	}

	email.NewSendEmail(toEmail, tpl.String(), data.SUBJECT, conf)
}

func SendForgotPassword(toEmail string, data dto.ForgotPasswordEmailRequestBody, conf *config.AppConfig) {
	t, err := templateRender.templates.ParseFS(views, "template/forgot_password.html")
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		panic(err)
	}

	email.NewSendEmail(toEmail, tpl.String(), data.SUBJECT, conf)
}

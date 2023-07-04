package config

import "os"

func initXendit(conf *AppConfig) {
	secret := os.Getenv("XENDIT_SECRET_KEY")
	successUrl := os.Getenv("XENDIT_SUCCESS_URL")

	conf.Xendit.SecretKey = secret
	conf.Xendit.SuccessUrl = successUrl
}

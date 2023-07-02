package config

import (
	"github.com/joho/godotenv"
)

// struct for the config file
type AppConfig struct {
	App struct {
		Env string
	}
	Fiber struct {
		Port string
	}
	Mysql struct {
		Host string
		Port string
		User string
		Pass string
		Name string
		Ssl  string
	}
	SendGrid struct {
		MailSenderName string
		MailSenderKey  string
	}
	Jwt struct {
		Secret string
	}
	Cloudinary struct {
		CloudName string
		ApiKey    string
		ApiSecret string
	}
}

var appConfig *AppConfig

func NewAppConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	if appConfig == nil {
		appConfig = &AppConfig{}

		initApp(appConfig)
		initFiber(appConfig)
		initMysql(appConfig)
		initSendGrid(appConfig)
		initJwt(appConfig)
		initCloudinary(appConfig)
	}

	return appConfig
}

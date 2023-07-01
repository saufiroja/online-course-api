package config

import "os"

func initJwt(conf *AppConfig) {
	secret := os.Getenv("JWT_SECRET")

	conf.Jwt.Secret = secret
}

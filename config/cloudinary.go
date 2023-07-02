package config

import "os"

func initCloudinary(conf *AppConfig) {
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	conf.Cloudinary.CloudName = cloudName
	conf.Cloudinary.ApiKey = apiKey
	conf.Cloudinary.ApiSecret = apiSecret
}

package cloudinary

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/saufiroja/online-course-api/config"
)

func NewCloudinary(conf *config.AppConfig) *cloudinary.Cloudinary {
	cloudName := conf.Cloudinary.CloudName
	apiKey := conf.Cloudinary.ApiKey
	apiSecret := conf.Cloudinary.ApiSecret

	cld, err := cloudinary.NewFromParams(
		cloudName,
		apiKey,
		apiSecret,
	)
	if err != nil {
		panic(err)
	}

	cld.Config.URL.Secure = true

	return cld
}

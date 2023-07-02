package utils

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
	"github.com/saufiroja/online-course-api/config"
	"github.com/saufiroja/online-course-api/infrastructure/cloudinary"
)

func UploadImage(file multipart.FileHeader, conf *config.AppConfig) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cld := cloudinary.NewCloudinary(conf)

	buf := new(bytes.Buffer)

	src, err := file.Open()
	if err != nil {
		return nil, err
	}

	defer src.Close()

	_, err = io.Copy(buf, src)
	if err != nil {
		return nil, err
	}

	res, err := cld.Upload.Upload(ctx, buf, uploader.UploadParams{
		Folder:   "online-course",
		PublicID: uuid.New().String(),
	})
	if err != nil {
		return nil, err
	}

	return &res.URL, nil
}

func DeleteImage(file string, conf *config.AppConfig) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cld := cloudinary.NewCloudinary(conf)

	// get image from cloudinary
	files := filepath.Base(file)

	str := files[:len(files)-len(filepath.Ext(files))]

	res, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: str,
	})
	if err != nil {
		return "", err
	}

	return res.Result, nil
}

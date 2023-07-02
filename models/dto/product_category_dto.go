package dto

import "mime/multipart"

type ProductCategoryRequestBody struct {
	Name      string                `json:"name" form:"name"`
	Image     *multipart.FileHeader `json:"image" form:"image"`
	CreatedBy *int64
	UpdatedBy *int64
}

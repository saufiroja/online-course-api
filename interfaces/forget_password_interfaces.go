package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/models/entity"
)

type ForgetPasswordRepository interface {
	InsertForgetPassword(input *entity.ForgotPassword) (*entity.ForgotPassword, error)
	FindOneByCode(code string) (*entity.ForgotPassword, error)
	UpdateForgetPassword(input *entity.ForgotPassword) (*entity.ForgotPassword, error)
}

type ForgetPasswordService interface {
	InsertForgetPassword(input *dto.ForgotPasswordRequestBody) (*entity.ForgotPassword, error)
	UpdateForgetPassword(input *dto.ForgotPasswordUpdateRequestBody) (*entity.ForgotPassword, error)
}

type ForgetPasswordHandler interface {
	InsertForgetPassword(c *fiber.Ctx) error
	UpdateForgetPassword(c *fiber.Ctx) error
}

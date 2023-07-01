package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/models/entity"
)

type AdminRepository interface {
	FindAllAdmin(offset, limit int) ([]entity.Admin, error)
	FindOneAdminByID(id int) (*entity.Admin, error)
	FindOneAdminByEmail(email string) (*entity.Admin, error)
	InsertAdmin(admin entity.Admin) (*entity.Admin, error)
	UpdateAdmin(admin entity.Admin) (*entity.Admin, error)
	DeleteAdmin(admin entity.Admin) error
	CountAdmin() (int64, error)
}

type AdminService interface {
	FindAllAdmin(offset, limit int) ([]entity.Admin, error)
	FindOneAdminByID(id int) (*entity.Admin, error)
	FindOneAdminByEmail(email string) (*entity.Admin, error)
	InsertAdmin(admin dto.AdminRequestBody) (*entity.Admin, error)
	UpdateAdmin(id int, admin dto.AdminRequestBody) (*entity.Admin, error)
	DeleteAdmin(id int) error
	CountAdmin() (int64, error)
}

type AdminHandler interface {
	FindAllAdmin(c *fiber.Ctx) error
	FindOneAdminByID(c *fiber.Ctx) error
	InsertAdmin(c *fiber.Ctx) error
	UpdateAdmin(c *fiber.Ctx) error
	DeleteAdmin(c *fiber.Ctx) error
}

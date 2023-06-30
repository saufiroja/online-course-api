package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/models/entity"
)

type UserRepository interface {
	FindAllUser(offset int, limit int) ([]entity.User, error)
	FindUserById(id int) (*entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
	InsertUser(user *entity.User) (*entity.User, error)
	FindUserByCodeVerified(codeVerified string) (*entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
	DeleteUser(user *entity.User) error
	CountUser() (int64, error)
}

type UserService interface {
	FindAllUser(offset int, limit int) ([]entity.User, error)
	FindUserById(id int) (*entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
	InsertUser(input *dto.UserRequestBody) (*entity.User, error)
	FindUserByCodeVerified(codeVerified string) (*entity.User, error)
	UpdateUser(id int, input *dto.UserUpdateRequestBody) (*entity.User, error)
	DeleteUser(id int) error
	CountUser() (int64, error)
}

type UserHandler interface {
	FindAllUser(c *fiber.Ctx) error
	FindUserById(c *fiber.Ctx) error
	InsertUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

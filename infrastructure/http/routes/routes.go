package routes

import (
	"github.com/saufiroja/online-course-api/config"
	"github.com/saufiroja/online-course-api/infrastructure/database"
	adminHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/admin"
	registerHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/register"
	userHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/user"
	"github.com/saufiroja/online-course-api/interfaces"
	adminRepo "github.com/saufiroja/online-course-api/repository/mysql/admin"
	userRepo "github.com/saufiroja/online-course-api/repository/mysql/user"
	adminSvc "github.com/saufiroja/online-course-api/service/admin"
	registerSvc "github.com/saufiroja/online-course-api/service/register"
	userSvc "github.com/saufiroja/online-course-api/service/user"

	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	// every new handler
	userHandler     interfaces.UserHandler
	registerHandler interfaces.RegisterHandler
	adminHandler    interfaces.AdminHandler
}

func NewRoutes(
	userHandler interfaces.UserHandler,
	registerHandler interfaces.RegisterHandler,
	adminHandler interfaces.AdminHandler,
) *Routes {
	return &Routes{
		userHandler:     userHandler,
		registerHandler: registerHandler,
		adminHandler:    adminHandler,
	}
}

func (r *Routes) initRoutes(app *fiber.App) {
	user := app.Group("/api/v1")
	// user
	user.Post("/users", r.userHandler.InsertUser)
	user.Get("/users", r.userHandler.FindAllUser)
	user.Get("/users/:id", r.userHandler.FindUserById)
	user.Patch("/users/:id", r.userHandler.UpdateUser)
	user.Delete("/users/:id", r.userHandler.DeleteUser)

	// register
	user.Post("/users/register", r.registerHandler.Register)

	// admin
	admin := app.Group("/api/v1")
	admin.Post("/admins", r.adminHandler.InsertAdmin)
	admin.Get("/admins", r.adminHandler.FindAllAdmin)
	admin.Get("/admins/:id", r.adminHandler.FindOneAdminByID)
	admin.Patch("/admins/:id", r.adminHandler.UpdateAdmin)
	admin.Delete("/admins/:id", r.adminHandler.DeleteAdmin)
}

func Initilized(app *fiber.App, conf *config.AppConfig) {
	db := database.NewMysql(conf)
	// repository
	user := userRepo.NewUserRepository(db)
	admin := adminRepo.NewAdminRepository(db)

	// service
	userSvc := userSvc.NewUserService(user)
	registerSvc := registerSvc.NewRegisterService(userSvc, conf)
	adminSvc := adminSvc.NewAdminService(admin)

	// handler
	userHandler := userHndlr.NewUserHandler(userSvc)
	registerHandler := registerHndlr.NewRegisterHandler(registerSvc)
	adminHandler := adminHndlr.NewAdminHandler(adminSvc)

	routes := NewRoutes(
		userHandler,
		registerHandler,
		adminHandler,
	)

	routes.initRoutes(app)
}

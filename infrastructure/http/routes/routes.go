package routes

import (
	"github.com/saufiroja/online-course-api/config"
	"github.com/saufiroja/online-course-api/infrastructure/database"
	registerHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/register"
	userHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/user"
	"github.com/saufiroja/online-course-api/interfaces"
	userRepo "github.com/saufiroja/online-course-api/repository/mysql/user"
	registerSvc "github.com/saufiroja/online-course-api/service/register"
	userSvc "github.com/saufiroja/online-course-api/service/user"

	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	// every new handler
	userHandler     interfaces.UserHandler
	registerHandler interfaces.RegisterHandler
}

func NewRoutes(
	userHandler interfaces.UserHandler,
	registerHandler interfaces.RegisterHandler,
) *Routes {
	return &Routes{
		userHandler:     userHandler,
		registerHandler: registerHandler,
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
}

func Initilized(app *fiber.App, conf *config.AppConfig) {
	db := database.NewMysql(conf)
	// repository
	user := userRepo.NewUserRepository(db)

	// service
	userSvc := userSvc.NewUserService(user)
	registerSvc := registerSvc.NewRegisterService(userSvc, conf)

	// handler
	userHandler := userHndlr.NewUserHandler(userSvc)
	registerHandler := registerHndlr.NewRegisterHandler(registerSvc)

	routes := NewRoutes(
		userHandler,
		registerHandler,
	)

	routes.initRoutes(app)
}

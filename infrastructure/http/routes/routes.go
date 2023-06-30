package routes

import (
	"github.com/saufiroja/online-course-api/config"
	"github.com/saufiroja/online-course-api/infrastructure/database"
	userHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/user"
	"github.com/saufiroja/online-course-api/interfaces"
	userRepo "github.com/saufiroja/online-course-api/repository/mysql/user"
	userSvc "github.com/saufiroja/online-course-api/service/user"

	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	// every new handler
	userHandler interfaces.UserHandler
}

func NewRoutes(userHandler interfaces.UserHandler) *Routes {
	return &Routes{
		userHandler: userHandler,
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
}

func Initilized(app *fiber.App, conf *config.AppConfig) {
	db := database.NewMysql(conf)
	// repository
	user := userRepo.NewUserRepository(db)

	// service
	userSvc := userSvc.NewUserService(user)

	// handler
	userHandler := userHndlr.NewUserHandler(userSvc)

	routes := NewRoutes(
		userHandler,
	)

	routes.initRoutes(app)
}

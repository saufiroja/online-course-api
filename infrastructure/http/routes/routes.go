package routes

import (
	"github.com/saufiroja/online-course-api/config"
	"github.com/saufiroja/online-course-api/infrastructure/database"
	adminHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/admin"
	forgotPasswordHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/forgetpassword"
	oauthHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/oauth"
	registerHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/register"
	userHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/user"
	"github.com/saufiroja/online-course-api/infrastructure/http/middlewares"
	"github.com/saufiroja/online-course-api/interfaces"
	adminRepo "github.com/saufiroja/online-course-api/repository/mysql/admin"
	forgotPasswordRepo "github.com/saufiroja/online-course-api/repository/mysql/forgetpassword"
	oauthAccessRepo "github.com/saufiroja/online-course-api/repository/mysql/oauth"
	oauthClientRepo "github.com/saufiroja/online-course-api/repository/mysql/oauth"
	oauthRefreshRepo "github.com/saufiroja/online-course-api/repository/mysql/oauth"
	userRepo "github.com/saufiroja/online-course-api/repository/mysql/user"
	adminSvc "github.com/saufiroja/online-course-api/service/admin"
	forgotPasswordSvc "github.com/saufiroja/online-course-api/service/forgetpassword"
	oauthSvc "github.com/saufiroja/online-course-api/service/oauth"
	registerSvc "github.com/saufiroja/online-course-api/service/register"
	userSvc "github.com/saufiroja/online-course-api/service/user"

	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	// every new handler
	userHandler           interfaces.UserHandler
	registerHandler       interfaces.RegisterHandler
	adminHandler          interfaces.AdminHandler
	oauthHandler          interfaces.OauthHandler
	forgotPasswordHandler interfaces.ForgetPasswordHandler
}

func NewRoutes(
	userHandler interfaces.UserHandler,
	registerHandler interfaces.RegisterHandler,
	adminHandler interfaces.AdminHandler,
	oauthHandler interfaces.OauthHandler,
	forgotPasswordHandler interfaces.ForgetPasswordHandler,
) *Routes {
	return &Routes{
		userHandler:           userHandler,
		registerHandler:       registerHandler,
		adminHandler:          adminHandler,
		oauthHandler:          oauthHandler,
		forgotPasswordHandler: forgotPasswordHandler,
	}
}

func (r *Routes) initRoutes(app *fiber.App) {
	users := app.Group("/api/v1/users")
	// user
	users.Use(middlewares.MiddlewaresUser, middlewares.MiddlewaresAdmin)
	users.Post("/", r.userHandler.InsertUser)
	users.Get("/", r.userHandler.FindAllUser)
	users.Get("/:id", r.userHandler.FindUserById)
	users.Patch("/:id", r.userHandler.UpdateUser)
	users.Delete("/:id", r.userHandler.DeleteUser)

	// register
	register := app.Group("/api/v1/register")
	register.Post("/", r.registerHandler.Register)

	// admin
	admin := app.Group("/api/v1/admins")
	admin.Post("/", r.adminHandler.InsertAdmin)
	admin.Get("/", r.adminHandler.FindAllAdmin)
	admin.Get("/:id", r.adminHandler.FindOneAdminByID)
	admin.Patch("/:id", r.adminHandler.UpdateAdmin)
	admin.Delete("/:id", r.adminHandler.DeleteAdmin)

	// oauth
	oauth := app.Group("/api/v1/oauth")
	oauth.Post("/login", r.oauthHandler.Login)
	oauth.Post("/refresh", r.oauthHandler.RefreshToken)

	// forgot password
	forgotPassword := app.Group("/api/v1/forgot-password")
	forgotPassword.Post("/", r.forgotPasswordHandler.InsertForgetPassword)
	forgotPassword.Put("/", r.forgotPasswordHandler.UpdateForgetPassword)
}

func Initilized(app *fiber.App, conf *config.AppConfig) {
	db := database.NewMysql(conf)
	// repository
	user := userRepo.NewUserRepository(db)
	admin := adminRepo.NewAdminRepository(db)
	oauthAccess := oauthAccessRepo.NewOauthAccessTokenRepository(db)
	oauthRefresh := oauthRefreshRepo.NewOauthRefreshTokenRepository(db)
	oauthClient := oauthClientRepo.NewOauthClientRepository(db)
	forgotPassword := forgotPasswordRepo.NewForgetPasswordRepository(db)

	// service
	userSvc := userSvc.NewUserService(user)
	registerSvc := registerSvc.NewRegisterService(userSvc, conf)
	adminSvc := adminSvc.NewAdminService(admin)
	oauthSvc := oauthSvc.NewOauthService(oauthAccess, oauthClient, oauthRefresh, adminSvc, userSvc, conf)
	forgotPasswordSvc := forgotPasswordSvc.NewForgetPasswordService(forgotPassword, userSvc, conf)

	// handler
	userHandler := userHndlr.NewUserHandler(userSvc)
	registerHandler := registerHndlr.NewRegisterHandler(registerSvc)
	adminHandler := adminHndlr.NewAdminHandler(adminSvc)
	oauthHandler := oauthHndlr.NewOauthHandler(oauthSvc)
	forgotPasswordHandler := forgotPasswordHndlr.NewForgetPasswordHandler(forgotPasswordSvc)

	routes := NewRoutes(
		userHandler,
		registerHandler,
		adminHandler,
		oauthHandler,
		forgotPasswordHandler,
	)

	routes.initRoutes(app)
}

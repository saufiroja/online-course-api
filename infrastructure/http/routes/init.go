package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/config"
	"github.com/saufiroja/online-course-api/infrastructure/database"
	adminHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/admin"
	forgotPasswordHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/forgetPassword"
	oauthHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/oauth"
	productCatgeoryHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/productCategory"
	registerHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/register"
	userHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/user"
	adminRepo "github.com/saufiroja/online-course-api/repository/mysql/admin"
	forgotPasswordRepo "github.com/saufiroja/online-course-api/repository/mysql/forgetPassword"
	oauthAccessRepo "github.com/saufiroja/online-course-api/repository/mysql/oauth"
	oauthClientRepo "github.com/saufiroja/online-course-api/repository/mysql/oauth"
	oauthRefreshRepo "github.com/saufiroja/online-course-api/repository/mysql/oauth"
	productCatgeoryRepo "github.com/saufiroja/online-course-api/repository/mysql/productCategory"
	userRepo "github.com/saufiroja/online-course-api/repository/mysql/user"
	adminSvc "github.com/saufiroja/online-course-api/service/admin"
	forgotPasswordSvc "github.com/saufiroja/online-course-api/service/forgetPassword"
	oauthSvc "github.com/saufiroja/online-course-api/service/oauth"
	productCatgeorySvc "github.com/saufiroja/online-course-api/service/productCategory"
	registerSvc "github.com/saufiroja/online-course-api/service/register"
	userSvc "github.com/saufiroja/online-course-api/service/user"
)

func Initilized(app *fiber.App, conf *config.AppConfig) {
	db := database.NewMysql(conf)
	// repository
	user := userRepo.NewUserRepository(db)
	admin := adminRepo.NewAdminRepository(db)
	oauthAccess := oauthAccessRepo.NewOauthAccessTokenRepository(db)
	oauthRefresh := oauthRefreshRepo.NewOauthRefreshTokenRepository(db)
	oauthClient := oauthClientRepo.NewOauthClientRepository(db)
	forgotPassword := forgotPasswordRepo.NewForgetPasswordRepository(db)
	productCategory := productCatgeoryRepo.NewProductCategoryRepository(db)

	// service
	userSvc := userSvc.NewUserService(user)
	registerSvc := registerSvc.NewRegisterService(userSvc, conf)
	adminSvc := adminSvc.NewAdminService(admin)
	oauthSvc := oauthSvc.NewOauthService(oauthAccess, oauthClient, oauthRefresh, adminSvc, userSvc, conf)
	forgotPasswordSvc := forgotPasswordSvc.NewForgetPasswordService(forgotPassword, userSvc, conf)
	productCategorySvc := productCatgeorySvc.NewProductCategoryService(productCategory, conf)

	// handler
	userHandler := userHndlr.NewUserHandler(userSvc)
	registerHandler := registerHndlr.NewRegisterHandler(registerSvc)
	adminHandler := adminHndlr.NewAdminHandler(adminSvc)
	oauthHandler := oauthHndlr.NewOauthHandler(oauthSvc)
	forgotPasswordHandler := forgotPasswordHndlr.NewForgetPasswordHandler(forgotPasswordSvc)
	productCategoryHandler := productCatgeoryHndlr.NewProductCategoryHandler(productCategorySvc)

	routes := NewRoutes(
		userHandler,
		registerHandler,
		adminHandler,
		oauthHandler,
		forgotPasswordHandler,
		productCategoryHandler,
	)

	routes.initRoutes(app)
}

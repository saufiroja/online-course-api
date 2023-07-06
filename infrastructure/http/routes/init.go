package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/config"
	"github.com/saufiroja/online-course-api/infrastructure/database"
	adminHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/admin"
	cartHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/cart"
	classRoomHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/classRoom"
	discountHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/discount"
	forgotPasswordHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/forgetPassword"
	oauthHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/oauth"
	orderHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/order"
	productHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/product"
	productCatgeoryHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/productCategory"
	registerHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/register"
	userHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/user"
	webhookHndlr "github.com/saufiroja/online-course-api/infrastructure/http/handler/webhook"
	adminRepo "github.com/saufiroja/online-course-api/repository/mysql/admin"
	cartRepo "github.com/saufiroja/online-course-api/repository/mysql/cart"
	classRoomRepo "github.com/saufiroja/online-course-api/repository/mysql/classRoom"
	discountRepo "github.com/saufiroja/online-course-api/repository/mysql/discount"
	forgotPasswordRepo "github.com/saufiroja/online-course-api/repository/mysql/forgetPassword"
	oauthAccessRepo "github.com/saufiroja/online-course-api/repository/mysql/oauth"
	oauthClientRepo "github.com/saufiroja/online-course-api/repository/mysql/oauth"
	oauthRefreshRepo "github.com/saufiroja/online-course-api/repository/mysql/oauth"
	orderRepo "github.com/saufiroja/online-course-api/repository/mysql/order"
	orderDetailRepo "github.com/saufiroja/online-course-api/repository/mysql/orderDetail"
	productRepo "github.com/saufiroja/online-course-api/repository/mysql/product"
	productCatgeoryRepo "github.com/saufiroja/online-course-api/repository/mysql/productCategory"
	userRepo "github.com/saufiroja/online-course-api/repository/mysql/user"
	adminSvc "github.com/saufiroja/online-course-api/service/admin"
	cartSvc "github.com/saufiroja/online-course-api/service/cart"
	classRoomSvc "github.com/saufiroja/online-course-api/service/classRoom"
	discountSvc "github.com/saufiroja/online-course-api/service/discount"
	forgotPasswordSvc "github.com/saufiroja/online-course-api/service/forgetPassword"
	oauthSvc "github.com/saufiroja/online-course-api/service/oauth"
	orderSvc "github.com/saufiroja/online-course-api/service/order"
	orderDetailSvc "github.com/saufiroja/online-course-api/service/orderDetail"
	paymentSvc "github.com/saufiroja/online-course-api/service/payment"
	productSvc "github.com/saufiroja/online-course-api/service/product"
	productCatgeorySvc "github.com/saufiroja/online-course-api/service/productCategory"
	registerSvc "github.com/saufiroja/online-course-api/service/register"
	userSvc "github.com/saufiroja/online-course-api/service/user"
	webhookSvc "github.com/saufiroja/online-course-api/service/webhook"
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
	product := productRepo.NewProductRepository(db)
	discount := discountRepo.NewDiscountRepository(db)
	cart := cartRepo.NewCartRepository(db)
	order := orderRepo.NewOrderRepository(db)
	orderDetail := orderDetailRepo.NewOrderDetailRepository(db)
	classRoom := classRoomRepo.NewClassRoomRepository(db)

	// service
	userSvc := userSvc.NewUserService(user)
	registerSvc := registerSvc.NewRegisterService(userSvc, conf)
	adminSvc := adminSvc.NewAdminService(admin)
	oauthSvc := oauthSvc.NewOauthService(oauthAccess, oauthClient, oauthRefresh, adminSvc, userSvc, conf)
	forgotPasswordSvc := forgotPasswordSvc.NewForgetPasswordService(forgotPassword, userSvc, conf)
	productCategorySvc := productCatgeorySvc.NewProductCategoryService(productCategory, conf)
	productSvc := productSvc.NewProductService(product, conf)
	discountSvc := discountSvc.NewDiscountService(discount)
	cartSvc := cartSvc.NewCartService(cart)
	orderDetailSvc := orderDetailSvc.NewOrderDetailService(orderDetail)
	paymentSvc := paymentSvc.NewPaymentService(conf)
	orderSvc := orderSvc.NewOrderService(order, cartSvc, discountSvc, productSvc, orderDetailSvc, paymentSvc)
	classRoomSvc := classRoomSvc.NewClassRoomService(classRoom)
	webhookSvc := webhookSvc.NewWebhookService(orderSvc, classRoomSvc, conf)

	// handler
	userHandler := userHndlr.NewUserHandler(userSvc)
	registerHandler := registerHndlr.NewRegisterHandler(registerSvc)
	adminHandler := adminHndlr.NewAdminHandler(adminSvc)
	oauthHandler := oauthHndlr.NewOauthHandler(oauthSvc)
	forgotPasswordHandler := forgotPasswordHndlr.NewForgetPasswordHandler(forgotPasswordSvc)
	productCategoryHandler := productCatgeoryHndlr.NewProductCategoryHandler(productCategorySvc)
	productHandler := productHndlr.NewProductHandler(productSvc)
	discountHandler := discountHndlr.NewDiscountHandler(discountSvc)
	cartHandler := cartHndlr.NewCartHandler(cartSvc)
	orderHandler := orderHndlr.NewOrderHandler(orderSvc)
	classRoomHandler := classRoomHndlr.NewClassRoomHandler(classRoomSvc)
	webhookHandler := webhookHndlr.NewWebhookHandler(webhookSvc)

	routes := NewRoutes(
		userHandler,
		registerHandler,
		adminHandler,
		oauthHandler,
		forgotPasswordHandler,
		productCategoryHandler,
		productHandler,
		discountHandler,
		cartHandler,
		orderHandler,
		classRoomHandler,
		webhookHandler,
	)

	routes.initRoutes(app)
}

package routes

import (
	"github.com/saufiroja/online-course-api/infrastructure/http/middlewares"
	"github.com/saufiroja/online-course-api/interfaces"

	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	// every new handler
	userHandler            interfaces.UserHandler
	registerHandler        interfaces.RegisterHandler
	adminHandler           interfaces.AdminHandler
	oauthHandler           interfaces.OauthHandler
	forgotPasswordHandler  interfaces.ForgetPasswordHandler
	productCategoryHandler interfaces.ProductCategoryHandler
	productHandler         interfaces.ProductHandler
	discountHandler        interfaces.DiscountHandler
	cartHandler            interfaces.CartHandler
}

func NewRoutes(
	userHandler interfaces.UserHandler,
	registerHandler interfaces.RegisterHandler,
	adminHandler interfaces.AdminHandler,
	oauthHandler interfaces.OauthHandler,
	forgotPasswordHandler interfaces.ForgetPasswordHandler,
	productCategoryHandler interfaces.ProductCategoryHandler,
	productHandler interfaces.ProductHandler,
	discountHandler interfaces.DiscountHandler,
	cartHandler interfaces.CartHandler,
) *Routes {
	return &Routes{
		userHandler:            userHandler,
		registerHandler:        registerHandler,
		adminHandler:           adminHandler,
		oauthHandler:           oauthHandler,
		forgotPasswordHandler:  forgotPasswordHandler,
		productCategoryHandler: productCategoryHandler,
		productHandler:         productHandler,
		discountHandler:        discountHandler,
		cartHandler:            cartHandler,
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
	admin.Use(middlewares.MiddlewaresUser, middlewares.MiddlewaresAdmin)
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

	// product category
	productCategory := app.Group("/api/v1/product-categories")
	productCategory.Use(middlewares.MiddlewaresUser, middlewares.MiddlewaresAdmin)
	productCategory.Post("/", r.productCategoryHandler.InsertProductCategory)
	productCategory.Get("/", r.productCategoryHandler.FindAllProductCategory)
	productCategory.Get("/:id", r.productCategoryHandler.FindProductCategoryByID)
	productCategory.Patch("/:id", r.productCategoryHandler.UpdateProductCategory)
	productCategory.Delete("/:id", r.productCategoryHandler.DeleteProductCategory)

	// product
	product := app.Group("/api/v1/products")
	product.Get("/", r.productHandler.FindAllProduct)
	product.Get("/:id", r.productHandler.FindProductByID)
	product.Use(middlewares.MiddlewaresUser, middlewares.MiddlewaresAdmin)
	product.Post("/", r.productHandler.InsertProduct)
	product.Patch("/:id", r.productHandler.UpdateProductByID)
	product.Delete("/:id", r.productHandler.DeleteProductByID)

	// discount
	discount := app.Group("/api/v1/discounts")
	discount.Use(middlewares.MiddlewaresUser, middlewares.MiddlewaresAdmin)
	discount.Post("/", r.discountHandler.InsertDiscount)
	discount.Get("/", r.discountHandler.FindAllDiscount)
	discount.Get("/:id", r.discountHandler.FindDiscountById)
	discount.Get("/:code/code", r.discountHandler.FindDiscountByCode)
	discount.Patch("/:id", r.discountHandler.UpdateDiscountById)
	discount.Patch("/:id/remaining-quantity", r.discountHandler.UpdateRemainingDiscount)
	discount.Delete("/:id", r.discountHandler.DeleteDiscountById)

	// cart
	cart := app.Group("/api/v1/carts")
	cart.Use(middlewares.MiddlewaresUser)
	cart.Post("/", r.cartHandler.InsertCart)
	cart.Get("/", r.cartHandler.FindAllCartByUserId)
	cart.Patch("/:id", r.cartHandler.UpdateCart)
	cart.Delete("/:id", r.cartHandler.DeleteCart)
}

package routes

import (
	"github.com/saufiroja/online-course-api/config"

	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	// every new handler

}

func NewRoutes() *Routes {
	return &Routes{}
}

func (r *Routes) initRoutes(app *fiber.App) {

}

func Initilized(app *fiber.App, conf *config.AppConfig) {

	// repository

	// service

	// controllers

	routes := NewRoutes()

	routes.initRoutes(app)
}

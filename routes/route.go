package routes

import (
	"onestep/controller"
	"onestep/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	// Authenticated Routes
	app.Use(middleware.IsAuthenticated)

	app.Get("/api/user", controller.User)
	app.Post("/api/logout", controller.Logout)

}

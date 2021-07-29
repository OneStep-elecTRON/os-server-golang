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

	// User CRUD
	app.Get("/api/users", controller.AllUsers)
	app.Post("/api/user", controller.CreateUser)
	app.Get("/api/user/:id", controller.GetUser)
	app.Put("/api/user/:id", controller.UpdateUser)
	app.Delete("/api/user/:id", controller.DeleteUser)

	// Roles CRUD
	app.Get("/api/roles", controller.AllRoles)
	app.Post("/api/role", controller.CreateRole)
	app.Get("/api/role/:id", controller.GetRole)
	app.Put("/api/role/:id", controller.UpdateRole)
	app.Delete("/api/role/:id", controller.DeleteRole)

}

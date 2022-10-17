package routes

import (
	"github.com/Justincletus003/auth-system/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Get("/api/user", controller.User)
	app.Post("/api/logout", controller.Logout)
}
package routes

import (
	"rbac-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func Auth(app *fiber.App) {
	app.Get("/", controllers.TestCall)
	app.Post("/login", controllers.Login)
}

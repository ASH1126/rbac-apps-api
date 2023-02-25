package routes

import (
	controllers "rbac-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func Auth(app *fiber.App) {
	app.Get("/", controllers.TestCall)
	app.Post("/login", controllers.Login)
}

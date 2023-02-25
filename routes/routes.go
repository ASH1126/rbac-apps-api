package routes

import (
	"rbac-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func Auth(app *fiber.App) {
	app.Get("/", handlers.TestCall)
	app.Post("/login", handlers.Login)
}

package main

import (
	"fmt"
	"log"
	"rbac-api/config"
	"rbac-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	config := config.LoadConfig(".")

	app := fiber.New(fiber.Config{
		Immutable: true,
	})

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Accept,Authorization,Content-Type,X-CSRF-Token",
		ExposeHeaders:    "Link",
		AllowCredentials: true,
		MaxAge:           300,
	}))

	app.Use(func(c *fiber.Ctx) error {
		// Set some security headers:
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Download-Options", "noopen")
		c.Set("Strict-Transport-Security", "max-age=5184000")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-DNS-Prefetch-Control", "off")

		// Go to next middleware:
		return c.Next()
	})

	signingKey := []byte("ashmind26")
	configjwt := jwtware.Config{
		TokenLookup:  "header:Authorization",
		ErrorHandler: app.ErrorHandler,
		SigningKey:   signingKey,
	}
	app2 := app.Group("/v1")
	app2.Use(jwtware.New(configjwt))

	// Routing
	routes.Auth(app)

	host := fmt.Sprintf(":%d", config.ServerPort)
	log.Fatal(app.Listen(host))
}

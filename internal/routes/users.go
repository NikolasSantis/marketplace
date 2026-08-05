package routes

import (
	"marketplace/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App) {
	user := app.Group("/user")

	user.Post("/", handlers.CreateUser())
}

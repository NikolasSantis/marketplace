package routes

import (
	"marketplace/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App) {
	user := app.Group("/user")

	user.Get("/", handlers.GetUser())
	user.Post("/", handlers.CreateUser())
	user.Put("/", handlers.UpdateUser())
	user.Delete("/", handlers.DeleteUser())
}

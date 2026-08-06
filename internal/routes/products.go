package routes

import (
	"marketplace/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterProductRoutes(app *fiber.App) {
	product := app.Group("product")

	product.Get("/", handlers.GetProduct())
	product.Post("/", handlers.CreateProduct())
	product.Put("/", handlers.UpdateProduct())
	product.Delete("/", handlers.DeleteProduct())
}

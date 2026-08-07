package routes

import (
	"marketplace/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterOrderRoutes(app *fiber.App) {
	product := app.Group("orders")

	product.Get("/", handlers.GetOrder())
	product.Post("/", handlers.CreateOrder())
	product.Patch("/", handlers.UpdateOrder())
	product.Delete("/", handlers.DeleteOrder())
}

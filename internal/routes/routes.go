package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	RegisterUserRoutes(app)
	RegisterProductRoutes(app)
}

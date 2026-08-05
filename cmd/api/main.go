package main

import (
	"marketplace/internal/database"
	"marketplace/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()

	app := fiber.New()

	routes.RegisterRoutes(app)

	app.Listen(":3000")
}

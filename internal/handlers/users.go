package handlers

import (
	"marketplace/internal/database"
	"marketplace/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var user models.User

		// return c.Send(c.Body())
		// Faz o "bind" do JSON que vem do Vue.js para a struct
		if err := c.BodyParser(&user); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Dados inválidos"})
		}

		now := time.Now().UTC()

		user.Created_at = now
		user.Updated_at = now

		// return c.JSON(user)

		collection := database.Collection("users")

		_, err := collection.InsertOne(c.Context(), user)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Erro ao criar usuário",
			})
		}

		return c.Status(201).JSON(user)
	}
}

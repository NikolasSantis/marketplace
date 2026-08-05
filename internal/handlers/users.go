package handlers

import (
	"context"
	"marketplace/internal/database"
	"marketplace/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.Query("email")

		if email == "" {
			return c.Status(400).JSON(fiber.Map{})
		}

		collection := database.GetCollection("users")

		ctx := context.Background()

		var user models.User

		err := collection.FindOne(ctx, bson.M{
			"email": email,
		}).Decode(&user)

		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "usuário não encontrado",
			})
		}

		return c.JSON(user)
	}
}

func CreateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var user models.User

		// Faz o "bind" do JSON que vem do Vue.js para a struct
		if err := c.BodyParser(&user); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Dados inválidos"})
		}

		now := time.Now().UTC()

		user.Created_at = now
		user.Updated_at = now

		collection := database.GetCollection("users")

		_, err := collection.InsertOne(c.Context(), user)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Erro ao criar usuário",
			})
		}

		return c.Status(201).JSON(user)
	}
}

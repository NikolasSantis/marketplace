package handlers

import (
	"context"
	"maps"
	"marketplace/internal/database"
	"marketplace/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetOrder() fiber.Handler {
	return func(c *fiber.Ctx) error {

		userIdHex := c.Query("user_id")

		if userIdHex == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Parâmetro necessário: user_id",
			})
		}

		userId, err := bson.ObjectIDFromHex(userIdHex)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "ID de usuário inválido",
			})
		}

		orderIdHex := c.Query("order_id")

		collection := database.GetCollection("orders")
		ctx := context.Background()

		if orderIdHex == "" {

			var orders []models.Order

			cursor, err := collection.Find(ctx, bson.M{
				"user_id": userId,
			})

			if err != nil {
				return c.Status(500).JSON(fiber.Map{
					"error": "Erro ao buscar pedidos",
				})
			}
			defer cursor.Close(ctx)

			if err := cursor.All(ctx, &orders); err != nil {
				return c.Status(500).JSON(fiber.Map{
					"error": "Erro ao decodificar pedidos",
				})
			}

			return c.Status(200).JSON(orders)
		}

		orderId, err := bson.ObjectIDFromHex(orderIdHex)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "ID de pedido inválido",
			})
		}

		var order models.Order

		err = collection.FindOne(ctx, bson.M{
			"_id":     orderId,
			"user_id": userId,
		}).Decode(&order)

		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "Pedido não encontrado",
			})
		}

		return c.Status(200).JSON(order)
	}
}

func CreateOrder() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var order models.Order

		if err := c.BodyParser(&order); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Dados inválidos",
			})
		}

		now := time.Now().UTC()

		order.Created_at = now
		order.Updated_at = now

		collection := database.GetCollection("orders")

		_, err := collection.InsertOne(c.Context(), order)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Erro ao criar pedido",
				"err":   err.Error(),
			})
		}

		return c.Status(201).JSON(order)
	}
}

func UpdateOrder() fiber.Handler {
	return func(c *fiber.Ctx) error {

		userIdHex := c.Query("user_id")
		orderIdHex := c.Query("order_id")

		if userIdHex == "" || orderIdHex == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Parâmetros faltando",
			})
		}

		userId, err := bson.ObjectIDFromHex(userIdHex)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "ID de usuário inválido",
			})
		}

		orderId, err := bson.ObjectIDFromHex(orderIdHex)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "ID de pedido inválido",
			})
		}

		var dataToUpdate map[string]any

		if err := c.BodyParser(&dataToUpdate); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Dados inválidos",
			})
		}

		dataToSet := bson.M{}

		maps.Copy(dataToSet, dataToUpdate)

		dataToSet["updated_at"] = time.Now().UTC()

		collection := database.GetCollection("orders")
		ctx := context.Background()

		updateResult, err := collection.UpdateOne(
			ctx,
			bson.M{
				"_id":     orderId,
				"user_id": userId,
			},
			bson.M{
				"$set": dataToSet,
			},
		)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if updateResult.MatchedCount == 0 {
			return c.Status(404).JSON(fiber.Map{
				"error": "Pedido não encontrado",
			})
		}

		if updateResult.ModifiedCount == 0 {
			return c.Status(200).JSON(fiber.Map{
				"message": "Nenhum campo alterado",
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "Pedido atualizado",
		})
	}
}

func DeleteOrder() fiber.Handler {
	return func(c *fiber.Ctx) error {

		userIdHex := c.Query("user_id")
		orderIdHex := c.Query("order_id")

		if userIdHex == "" || orderIdHex == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Parâmetros faltando",
			})
		}

		userId, err := bson.ObjectIDFromHex(userIdHex)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "ID de usuário inválido",
			})
		}

		orderId, err := bson.ObjectIDFromHex(orderIdHex)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "ID de pedido inválido",
			})
		}

		ctx := context.Background()

		collection := database.GetCollection("orders")

		deleteResult, err := collection.DeleteOne(
			ctx,
			bson.M{
				"_id":     orderId,
				"user_id": userId,
			},
		)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Erro ao deletar pedido",
			})
		}

		if deleteResult.DeletedCount == 0 {
			return c.Status(404).JSON(fiber.Map{
				"error": "Pedido não encontrado",
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "Pedido deletado",
		})
	}
}

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

func GetProduct() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userIdHex := c.Query("user_id")

		if userIdHex == "" {
			return c.Status(400).JSON(fiber.Map{
				"erro": "Parâmetro necessário: user_id",
			})
		}

		userId, err := bson.ObjectIDFromHex(userIdHex)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"erro": "Verifique os dados de usuário",
			})
		}

		productIdHex := c.Query("product_id")

		collection := database.GetCollection("products")
		ctx := context.Background()

		if productIdHex == "" {

			var products []models.Product

			cursor, err := collection.Find(ctx, bson.M{
				"user_id": userId,
			})

			if err != nil {
				return c.Status(500).JSON(fiber.Map{
					"error":       "Erro ao buscar produtos",
					"description": err.Error(),
				})
			}
			defer cursor.Close(ctx)

			if err := cursor.All(ctx, &products); err != nil {
				return c.Status(500).JSON(fiber.Map{
					"error": "Erro ao decodificar usuários",
				})
			}

			return c.Status(200).JSON(products)
		}

		var products models.Product

		idToGet, err := bson.ObjectIDFromHex(productIdHex)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "ID de produto inválido",
			})
		}

		findErr := collection.FindOne(ctx, bson.M{
			"_id": idToGet,
		}).Decode(&products)

		if findErr != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "Produto não encontrado",
			})
		}

		return c.Status(200).JSON(products)
	}
}

func CreateProduct() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var product models.Product

		if err := c.BodyParser(&product); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Dados inválidos"})
		}

		now := time.Now().UTC()

		product.Created_at, product.Updated_at = now, now

		collection := database.GetCollection("products")

		_, err := collection.InsertOne(c.Context(), product)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Erro ao criar Produto",
				"err":   err.Error(),
			})
		}

		return c.Status(200).JSON(product)
	}
}

func UpdateProduct() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userIdHex := c.Query("user_id")
		productIdHex := c.Query("product_id")

		if userIdHex == "" || productIdHex == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Parâmetros Faltando",
			})
		}

		userId, err := bson.ObjectIDFromHex(userIdHex)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{"erro": "Id de usuário em formato inválido"})
		}

		productId, err := bson.ObjectIDFromHex(productIdHex)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{"erro": "Id de produto em formato inválido"})
		}

		var dataToUpdate map[string]any

		if err := c.BodyParser(&dataToUpdate); err != nil {
			return c.Status(400).JSON(fiber.Map{"erro": "Dados de produtos inválidos"})
		}

		dataToSet := bson.M{}

		maps.Copy(dataToSet, dataToUpdate)

		dataToSet["updated_at"] = time.Now().UTC()

		collection := database.GetCollection("products")
		ctx := context.Background()

		updateProductResult, err := collection.UpdateOne(ctx,
			bson.M{
				"_id":     productId,
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

		if updateProductResult.MatchedCount == 0 {
			return c.Status(404).JSON(fiber.Map{
				"error": "Produto não encontrado",
			})
		}

		if updateProductResult.ModifiedCount == 0 {
			return c.Status(200).JSON(fiber.Map{
				"message": "Nenhum campo foi alterado",
			})
		}

		return c.Status(200).JSON(fiber.Map{"message": "Produto atualizado"})
	}
}

func DeleteProduct() fiber.Handler {
	return func(c *fiber.Ctx) error {
		productIdHex := c.Query("product_id")
		userIdHex := c.Query("user_id")

		if productIdHex == "" || userIdHex == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Parâmetros Faltando"})
		}

		userId, err := bson.ObjectIDFromHex(userIdHex)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Id de usuário no formato inválido"})
		}

		productId, err := bson.ObjectIDFromHex(productIdHex)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Id de produto com formato inválido"})
		}

		ctx := context.Background()
		collection := database.GetCollection("products")

		deleteResult, err := collection.DeleteOne(ctx,
			bson.M{
				"_id":     productId,
				"user_id": userId,
			},
		)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Erro ao deletar produto"})
		}

		if deleteResult.DeletedCount == 0 {
			return c.Status(404).JSON(fiber.Map{"error": "Produto não encontrado"})
		}

		return c.Status(200).JSON(fiber.Map{"message": "Produto deletado"})
	}
}

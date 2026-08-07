package models

import "go.mongodb.org/mongo-driver/v2/bson"

type OrderItem struct {
	ProductID bson.ObjectID `bson:"product_id" json:"product_id"`
	Name      string        `bson:"name" json:"name"`
	SKU       string        `bson:"sku" json:"sku"`
	Quantity  int64         `bson:"quantity" json:"quantity"`
	UnitPrice float64       `bson:"unit_price" json:"unit_price"`
	Subtotal  float64       `bson:"subtotal" json:"subtotal"`
}

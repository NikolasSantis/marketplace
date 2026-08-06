package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Product struct {
	ID         bson.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID     bson.ObjectID `bson:"user_id" json:"user_id"`
	Name       string        `bson:"name" json:"name"`
	SKU        string        `bson:"sku" json:"sku"`
	Price      float64       `bson:"price" json:"price"`
	Quantity   int64         `bson:"quantity" json:"quantity"`
	EAN        string        `bson:"ean, omitempty" json:"ean"`
	Height     float64       `bson:"height" json:"height"`
	Width      float64       `bson:"width" json:"width"`
	Depth      float64       `bson:"depth" json:"depth"`
	Weight     float64       `bson:"weight" json:"weight"`
	Created_at time.Time     `bson:"created_at" json:"created_at"`
	Updated_at time.Time     `bson:"updated_at" json:"updated_at"`
	Deleted_at time.Time     `bson:"deleted_at, omitempty" json:"deleted_at"`
}

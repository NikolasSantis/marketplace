package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Order struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      bson.ObjectID `bson:"user_id" json:"user_id"`
	BuyerEmail  string        `bson:"buyer_email" json:"buyer_email"`
	OrderDate   string        `bson:"order_date" json:"order_date"`
	OrderStatus string        `bson:"order_status" json:"order_status"`

	Items []OrderItem `bson:"items" json:"items"`

	Total      float64   `bson:"total" json:"total"`
	Created_at time.Time `bson:"created_at" json:"created_at"`
	Updated_at time.Time `bson:"updated_at" json:"updated_at"`
	DeletedAt  time.Time `bson:"deleted_at,omitempty" json:"deleted_at"`
}

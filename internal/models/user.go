package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID         bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Name       string        `bson:"name" json:"name"`
	Email      string        `bson:"email" json:"email"`
	Created_at time.Time     `bson:"created_at" json:"created_at"`
	Updated_at time.Time     `bson:"updated_at" json:"updated_at"`
	Deleted_at time.Time     `bson:"deleted_at,omitempty" json:"deleted_at"`
}

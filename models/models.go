package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoBSON struct {
	ID        primitive.ObjectID `bson:"id,omitempty"`
	Name      string             `bson:"name"`
	Completed bool               `bson:"completed"`
	CreatedAt time.Time          `bson:"createdat"`
}

type TodoJSON struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

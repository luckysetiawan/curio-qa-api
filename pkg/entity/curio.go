// Package entity stores all entity structs that the server uses.
package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

// Curio represents a curio with an ID, content, sender, and status.
type Curio struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Content string             `json:"content" bson:"content"`
	Sender  User               `json:"from" bson:"from"`
	Status  bool               `json:"status" bson:"status"`
}

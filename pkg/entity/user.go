// Package entity stores all entity structs that the server uses.
package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents an user with an ID, display name, username, password, and
// curios.
type User struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	DisplayName string             `json:"displayName" bson:"displayName"`
	Username    string             `json:"username" bson:"username"`
	Password    string             `json:"password,omitempty" bson:"password,omitempty"`
	Curios      []Curio            `json:"curios,omitempty" bson:"curios,omitempty"`
}

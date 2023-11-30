package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	DisplayName string             `json:"name" bson:"name"`
	UserName    string             `json:"username" bson:"username"`
	Password    string             `json:"password,omitempty" bson:"password,omitempty"`
	Curios      []Curio            `json:"curios" bson:"curios"`
}

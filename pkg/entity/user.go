package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	DisplayName string              `json:"name" bson:"name"`
	UserName    string              `json:"username" bson:"username"`
	Password    string              `json:"password,omitempty" bson:"password,omitempty"`
	Timestamp   primitive.Timestamp `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	Curios      []Curio             `json:"curios" bson:"curios"`
}

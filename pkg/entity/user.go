package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	DisplayName string             `json:"displayName" bson:"displayName"`
	UserName    string             `json:"userName" bson:"userName"`
	Password    string             `json:"password,omitempty" bson:"password,omitempty"`
	Curios      []Curio            `json:"curios" bson:"curios"`
}

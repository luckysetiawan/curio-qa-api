package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Curio struct {
	ID        primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	Content   string              `json:"content" bson:"content"`
	From      string              `json:"from" bson:"from"`
	Timestamp primitive.Timestamp `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	Status    bool                `json:"status" bson:"status"`
}

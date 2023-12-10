package util

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPrimitiveTimestamp() primitive.Timestamp {
	currentTime := uint32(time.Now().Unix())
	primitiveTimestamp := primitive.Timestamp{
		T: currentTime,
		I: 0,
	}

	return primitiveTimestamp
}

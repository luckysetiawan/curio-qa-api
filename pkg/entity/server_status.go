// Package entity stores all entity structs that the server uses.
package entity

// ServerStatus represents the server status describing mongo and redis statuses.
type ServerStatus struct {
	MongoStatus bool `json:"mongoStatus" bson:"mongoStatus"`
	RedisStatus bool `json:"redisStatus" bson:"redisStatus"`
}

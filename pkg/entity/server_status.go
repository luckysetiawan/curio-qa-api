package entity

type ServerStatus struct {
	MongoStatus bool `json:"mongoStatus" bson:"mongoStatus"`
	RedisStatus bool `json:"redisStatus" bson:"redisStatus"`
}

package entity

type ServerStatus struct {
	MongoStatus bool `json:"mongostatus" bson:"mongostatus"`
	RedisStatus bool `json:"redisstatus" bson:"redisstatus"`
}

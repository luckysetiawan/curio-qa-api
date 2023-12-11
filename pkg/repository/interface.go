package repository

import (
	"github.com/luckysetiawan/curio-qa-api/pkg/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IServerStatusRepository interface {
	GetServerStatus() entity.ServerStatus
}

type IUserRepository interface {
	CheckUsernameTaken(username string) bool
	Find(filter primitive.D, args ...*options.FindOneOptions) (entity.User, error)
	Insert(user entity.User) (interface{}, error)
}

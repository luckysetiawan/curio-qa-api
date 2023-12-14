package repository

import (
	"github.com/luckysetiawan/curio-qa-api/pkg/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ICurioRepository interface {
	Find(userID, curioID primitive.ObjectID) (entity.Curio, error)
	Insert(userID primitive.ObjectID, curio entity.Curio) error
	UpdateStatus(userID, curioID primitive.ObjectID, status bool) error
}

type IServerStatusRepository interface {
	GetServerStatus() entity.ServerStatus
}

type IUserRepository interface {
	GetLoginStatuses() ([]string, error)
	CheckUsernameTaken(username string) bool
	GetAll(filter primitive.M, args ...*options.FindOptions) ([]entity.User, error)
	Find(filter primitive.M, args ...*options.FindOneOptions) (entity.User, error)
	Insert(user entity.User) (interface{}, error)
	MarkLoginStatus(userID string, username string) error
	ClearLoginStatus(userID string) error
}

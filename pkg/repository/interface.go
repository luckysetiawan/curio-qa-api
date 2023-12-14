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
	GetLoginStatuses() ([]string, error)
	CheckUsernameTaken(username string) bool
	Find(filter primitive.D, args ...*options.FindOneOptions) (entity.User, error)
	Insert(user entity.User) (interface{}, error)
	MarkLoginStatus(userID string, username string) error
	ClearLoginStatus(userID string) error
}

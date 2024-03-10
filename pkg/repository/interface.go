// Package repository stores all database logic the server uses.
package repository

import (
	"github.com/luckysetiawan/curio-qa-api/pkg/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ICurioRepository defines methods that must be met for curio repository.
type ICurioRepository interface {
	Find(userID, curioID primitive.ObjectID) (entity.Curio, error)
	Insert(userID primitive.ObjectID, curio entity.Curio) error
	UpdateStatus(userID, curioID primitive.ObjectID, status bool) error
}

// IServerStatusRepository defines methods that must be met for server status
// repository.
type IServerStatusRepository interface {
	GetServerStatus() entity.ServerStatus
}

// IUserRepository defines methods that must be met for user repository.
type IUserRepository interface {
	GetLoginStatuses() ([]string, error)
	CheckUsernameTaken(username string) bool
	GetAll(filter primitive.M, args ...*options.FindOptions) ([]entity.User, error)
	Find(filter primitive.M, args ...*options.FindOneOptions) (entity.User, error)
	Insert(user entity.User) (interface{}, error)
	MarkLoginStatus(userID string, username string) error
	ClearLoginStatus(userID string) error
}

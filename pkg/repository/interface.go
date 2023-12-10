package repository

import "github.com/luckysetiawan/curio-qa-api/pkg/entity"

type IServerStatusRepository interface {
	GetServerStatus() entity.ServerStatus
}

type IUserRepository interface {
	Insert(user entity.User) (interface{}, error)
}

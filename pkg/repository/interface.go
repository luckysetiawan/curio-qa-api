package repository

import "github.com/luckysetiawan/curio-qa-api/pkg/entity"

type IServerStatusRepository interface {
	GetServerStatus() (serverStatus entity.ServerStatus)
}

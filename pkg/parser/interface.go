package parser

import (
	"net/http"

	"github.com/luckysetiawan/curio-qa-api/pkg/entity"
)

type IUserParser interface {
	ParseUserEntity(r *http.Request) (entity.User, error)
}

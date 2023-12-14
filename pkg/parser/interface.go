package parser

import (
	"net/http"

	"github.com/luckysetiawan/curio-qa-api/pkg/entity"
)

type ICurioParser interface {
	ParseCurioEntity(r *http.Request) (entity.Curio, error)
	ParseCurioID(r *http.Request) string
	ParseCurioReceiverUsername(r *http.Request) string
}

type IUserParser interface {
	ParseUserEntity(r *http.Request) (entity.User, error)
	ParseUsername(r *http.Request) string
}

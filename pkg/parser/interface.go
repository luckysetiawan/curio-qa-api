// Package parser stores all parsing logic the server uses.
package parser

import (
	"net/http"

	"github.com/luckysetiawan/curio-qa-api/pkg/entity"
)

// ICurioParser defines methods that must be met in curio Parser.
type ICurioParser interface {
	ParseCurioEntity(r *http.Request) (entity.Curio, error)
	ParseCurioID(r *http.Request) string
	ParseCurioReceiverUsername(r *http.Request) string
}

// IUserParser defines methods that must be met in user Parser.
type IUserParser interface {
	ParseUserEntity(r *http.Request) (entity.User, error)
	ParseUsername(r *http.Request) string
}

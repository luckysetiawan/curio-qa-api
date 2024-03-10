// Package parser stores all parsing logic the server uses.
package parser

import (
	"net/http"

	"github.com/luckysetiawan/curio-qa-api/internal/util"
	"github.com/luckysetiawan/curio-qa-api/pkg/entity"
)

// userParser stores parsing logic functions.
type userParser struct{}

// NewUserParser returns userParser struct.
func NewUserParser() IUserParser {
	return &userParser{}
}

// ParseUserEntity returns user entity from JSON data.
func (*userParser) ParseUserEntity(r *http.Request) (entity.User, error) {
	var (
		user entity.User
		err  error
	)

	err = util.ParseJSON(r, &user)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

// ParseUsername returns username from path parameter.
func (*userParser) ParseUsername(r *http.Request) string {
	ID := util.ParsePathParam(r, "username")

	return ID
}

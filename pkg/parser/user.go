package parser

import (
	"net/http"

	"github.com/luckysetiawan/curio-qa-api/internal/util"
	"github.com/luckysetiawan/curio-qa-api/pkg/entity"
)

type userParser struct{}

func NewUserParser() IUserParser {
	return &userParser{}
}

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

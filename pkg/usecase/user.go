package usecase

import (
	"net/http"

	"github.com/luckysetiawan/curio-qa-api/internal/webserver"
	"github.com/luckysetiawan/curio-qa-api/pkg/constant"
	"github.com/luckysetiawan/curio-qa-api/pkg/entity"
	"github.com/luckysetiawan/curio-qa-api/pkg/parser"
	"github.com/luckysetiawan/curio-qa-api/pkg/repository"
)

type userUseCase struct {
	parser        parser.IUserParser
	jsonPresenter webserver.IPresenterJSON
	repository    repository.IUserRepository
}

func NewUserUseCase(parser parser.IUserParser, jsonPresenter webserver.IPresenterJSON, repository repository.IUserRepository) *userUseCase {
	return &userUseCase{
		parser:        parser,
		jsonPresenter: jsonPresenter,
		repository:    repository,
	}
}

func (h *userUseCase) Insert(w http.ResponseWriter, r *http.Request) {
	var (
		user       entity.User
		insertedID interface{}
		err        error
	)

	user, err = h.parser.ParseUserEntity(r)
	if err != nil {
		h.jsonPresenter.SendError(w, constant.ErrorParsingMessage)
	}

	insertedID, err = h.repository.Insert(user)
	if err != nil {
		h.jsonPresenter.SendError(w, constant.ErrorGeneralMessage)
	}

	h.jsonPresenter.SendSuccess(w, insertedID)
}

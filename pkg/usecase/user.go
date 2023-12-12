package usecase

import (
	"net/http"

	"github.com/luckysetiawan/curio-qa-api/internal/constant"
	"github.com/luckysetiawan/curio-qa-api/internal/util"
	"github.com/luckysetiawan/curio-qa-api/internal/webserver"
	"github.com/luckysetiawan/curio-qa-api/pkg/parser"
	"github.com/luckysetiawan/curio-qa-api/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
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

func (u *userUseCase) Login(w http.ResponseWriter, r *http.Request) {
	user, err := u.parser.ParseUserEntity(r)
	if err != nil {
		u.jsonPresenter.SendError(w, constant.ErrorParsingMessage)
		return
	}

	filter := bson.D{{Key: "username", Value: user.Username}}

	data, err := u.repository.Find(filter)
	if err != nil {
		u.jsonPresenter.SendError(w, constant.ErrorGeneralMessage)
		return
	}

	err = util.ComparePassword(data.Password, user.Password)
	if err != nil {
		u.jsonPresenter.SendError(w, constant.ErrorUsernamePasswordMessage)
		return
	}

	webserver.GenerateToken(w, r, data.ID.String(), data.Username, constant.RegisteredUser)
	u.jsonPresenter.SendSuccess(w)
}

func (u *userUseCase) Logout(w http.ResponseWriter, r *http.Request) {
	webserver.ResetToken(w)
	u.jsonPresenter.SendSuccess(w)
}

func (u *userUseCase) Insert(w http.ResponseWriter, r *http.Request) {
	user, err := u.parser.ParseUserEntity(r)
	if err != nil {
		u.jsonPresenter.SendError(w, constant.ErrorParsingMessage)
		return
	}

	// check username availability
	usernameTaken := u.repository.CheckUsernameTaken(user.Username)
	if usernameTaken {
		u.jsonPresenter.SendError(w, constant.ErrorUsernameTakenMessage)
		return
	}

	insertedID, err := u.repository.Insert(user)
	if err != nil {
		u.jsonPresenter.SendError(w, constant.ErrorGeneralMessage)
		return
	}

	u.jsonPresenter.SendSuccess(w, insertedID)
}

// Package usecase stores all usecase logic the server uses.
package usecase

import (
	"net/http"

	"github.com/luckysetiawan/curio-qa-api/internal/constant"
	"github.com/luckysetiawan/curio-qa-api/internal/util"
	"github.com/luckysetiawan/curio-qa-api/internal/webserver"
	"github.com/luckysetiawan/curio-qa-api/pkg/parser"
	"github.com/luckysetiawan/curio-qa-api/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// curioUseCase stores parser, jsonPresenter, repository and user logic
// functions.
type userUseCase struct {
	parser        parser.IUserParser
	jsonPresenter webserver.IPresenterJSON
	repository    repository.IUserRepository
}

// NewUserUseCase returns userUseCase struct.
func NewUserUseCase(parser parser.IUserParser, jsonPresenter webserver.IPresenterJSON, repository repository.IUserRepository) *userUseCase {
	return &userUseCase{
		parser:        parser,
		jsonPresenter: jsonPresenter,
		repository:    repository,
	}
}

// GetAll gets all user data.
func (u *userUseCase) GetAll(w http.ResponseWriter, r *http.Request) {
	filter := bson.M{}
	args := options.Find().SetProjection(bson.M{"password": 0})

	users, err := u.repository.GetAll(filter, args)
	if err != nil {
		u.jsonPresenter.SendError(w, constant.ErrorGeneralMessage)
		return
	}

	u.jsonPresenter.SendSuccess(w, users)
}

// Find finds user data by username.
func (u *userUseCase) Find(w http.ResponseWriter, r *http.Request) {
	username := u.parser.ParseUsername(r)

	filter := bson.M{"username": username}
	args := options.FindOne().SetProjection(bson.M{"password": 0})

	user, err := u.repository.Find(filter, args)
	if err != nil {
		u.jsonPresenter.SendError(w, constant.ErrorGeneralMessage)
		return
	}

	u.jsonPresenter.SendSuccess(w, user)
}

// GetAllActiveUsers gets all current active users.
func (u *userUseCase) GetAllActiveUsers(w http.ResponseWriter, r *http.Request) {
	activeUserIDs, err := u.repository.GetLoginStatuses()
	if err != nil {
		u.jsonPresenter.SendError(w, constant.ErrorCacheMessage)
		return
	}

	u.jsonPresenter.SendSuccessWithCount(w, activeUserIDs, len(activeUserIDs))
}

// Login validates user credentials and sets the cookies.
func (u *userUseCase) Login(w http.ResponseWriter, r *http.Request) {
	user, err := u.parser.ParseUserEntity(r)
	if err != nil {
		u.jsonPresenter.SendError(w, constant.ErrorParsingMessage)
		return
	}

	filter := bson.M{"username": user.Username}

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

	err = u.repository.MarkLoginStatus(data.ID.Hex(), data.Username)
	if err != nil {
		u.jsonPresenter.SendError(w, constant.ErrorCacheMessage)
		return
	}

	webserver.GenerateToken(w, r, data.ID.Hex(), data.Username, constant.RegisteredUser)
	u.jsonPresenter.SendSuccess(w)
}

// Logout invalidates user cookies.
func (u *userUseCase) Logout(w http.ResponseWriter, r *http.Request) {
	userID, _ := webserver.GetDataFromCookies(r)

	err := u.repository.ClearLoginStatus(userID)
	if err != nil {
		u.jsonPresenter.SendError(w, constant.ErrorCacheMessage)
		return
	}

	webserver.ResetToken(w)
	u.jsonPresenter.SendSuccess(w)
}

// Insert inserts a new user.
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

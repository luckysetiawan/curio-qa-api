package usecase

import (
	"net/http"

	"github.com/luckysetiawan/curio-qa-api/internal/constant"
	"github.com/luckysetiawan/curio-qa-api/internal/webserver"
	"github.com/luckysetiawan/curio-qa-api/pkg/parser"
	"github.com/luckysetiawan/curio-qa-api/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type curioUseCase struct {
	parser         parser.ICurioParser
	jsonPresenter  webserver.IPresenterJSON
	repository     repository.ICurioRepository
	userRepository repository.IUserRepository
}

func NewCurioUseCase(parser parser.ICurioParser, jsonPresenter webserver.IPresenterJSON, repository repository.ICurioRepository, userRepository repository.IUserRepository) *curioUseCase {
	return &curioUseCase{
		parser:         parser,
		jsonPresenter:  jsonPresenter,
		repository:     repository,
		userRepository: userRepository,
	}
}

func (u *curioUseCase) Insert(w http.ResponseWriter, r *http.Request) {
	// Get data from front end
	curio, err := u.parser.ParseCurioEntity(r)
	if err != nil {
		u.jsonPresenter.SendError(w, constant.ErrorParsingMessage)
		return
	}
	_, senderUsername := webserver.GetDataFromCookies(r)

	receiverUsername := u.parser.ParseCurioReceiverUsername(r)
	if err != nil {
		u.jsonPresenter.SendError(w, constant.ErrorParsingMessage)
		return
	}

	// Get sender and receiver data from database
	filter := bson.M{"username": senderUsername}
	args := options.FindOne().SetProjection(bson.M{"_id": 1, "displayName": 1, "username": 1})

	sender, err := u.userRepository.Find(filter, args)
	if err != nil {
		u.jsonPresenter.SendError(w, constant.ErrorGeneralMessage)
		return
	}

	filter = bson.M{"username": receiverUsername}
	args = options.FindOne().SetProjection(bson.M{"_id": 1})

	receiver, err := u.userRepository.Find(filter, args)
	if err != nil {
		u.jsonPresenter.SendError(w, constant.ErrorGeneralMessage)
		return
	}

	// Fill the necessary values to curio
	curio.ID = primitive.NewObjectID()
	curio.Sender = sender

	// Insert curio
	err = u.repository.Insert(receiver.ID, curio)
	if err != nil {
		u.jsonPresenter.SendError(w, constant.ErrorGeneralMessage)
		return
	}

	u.jsonPresenter.SendSuccess(w)
}

func (u *curioUseCase) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	// Get data from front end
	stringCurioID := u.parser.ParseCurioID(r)
	if stringCurioID == "" {
		u.jsonPresenter.SendError(w, constant.ErrorParsingMessage)
		return
	}

	stringUserID, _ := webserver.GetDataFromCookies(r)

	userID, err := primitive.ObjectIDFromHex(stringUserID)
	if err != nil {
		u.jsonPresenter.SendError(w, constant.ErrorGeneralMessage)
		return
	}

	curioID, err := primitive.ObjectIDFromHex(stringCurioID)
	if err != nil {
		u.jsonPresenter.SendError(w, constant.ErrorGeneralMessage)
		return
	}

	// Get curio to check the current curio status
	curio, err := u.repository.Find(userID, curioID)
	if err != nil {
		u.jsonPresenter.SendError(w, constant.ErrorGeneralMessage)
		return
	}

	// Update curio status
	err = u.repository.UpdateStatus(userID, curioID, !curio.Status)
	if err != nil {
		u.jsonPresenter.SendError(w, constant.ErrorGeneralMessage)
		return
	}

	u.jsonPresenter.SendSuccess(w)
}

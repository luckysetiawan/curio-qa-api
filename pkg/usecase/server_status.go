package usecase

import (
	"net/http"

	"github.com/luckysetiawan/curio-qa-api/internal/webserver"
	"github.com/luckysetiawan/curio-qa-api/pkg/repository"
)

type serverStatusUseCase struct {
	jsonPresenter webserver.IPresenterJSON
	repository    repository.IServerStatusRepository
}

func NewServerStatusUseCase(jsonPresenter webserver.IPresenterJSON, repository repository.IServerStatusRepository) *serverStatusUseCase {
	return &serverStatusUseCase{
		jsonPresenter: jsonPresenter,
		repository:    repository,
	}
}

func (u *serverStatusUseCase) GetStatus(w http.ResponseWriter, r *http.Request) {
	result := u.repository.GetServerStatus()

	u.jsonPresenter.SendSuccess(w, result)
}

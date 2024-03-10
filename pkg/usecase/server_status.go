// Package usecase stores all usecase logic the server uses.
package usecase

import (
	"net/http"

	"github.com/luckysetiawan/curio-qa-api/internal/webserver"
	"github.com/luckysetiawan/curio-qa-api/pkg/repository"
)

// serverStatusUseCase stores jsonPresenter, repository, and server status logic
// functions.
type serverStatusUseCase struct {
	jsonPresenter webserver.IPresenterJSON
	repository    repository.IServerStatusRepository
}

// NewServerStatusUseCase returns serverStatusUseCase struct.
func NewServerStatusUseCase(jsonPresenter webserver.IPresenterJSON, repository repository.IServerStatusRepository) *serverStatusUseCase {
	return &serverStatusUseCase{
		jsonPresenter: jsonPresenter,
		repository:    repository,
	}
}

// GetStatus returns server status.
func (u *serverStatusUseCase) GetStatus(w http.ResponseWriter, r *http.Request) {
	result := u.repository.GetServerStatus()

	u.jsonPresenter.SendSuccess(w, result)
}

package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/luckysetiawan/curio-qa-api/internal/constant"
)

type jsonPresenter struct{}

func NewJsonPresenter() IPresenterJSON {
	return &jsonPresenter{}
}

func (*jsonPresenter) SendSuccess(w http.ResponseWriter, data ...interface{}) {
	w.Header().Set("Content-Type", "application/json")

	var response = BaseResponse{
		Status:  http.StatusOK,
		Message: constant.SuccessGeneralMessage,
	}

	if len(data) > 0 {
		response.Data = data[0]
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
	}
}

func (*jsonPresenter) SendError(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")

	response := BaseResponse{
		Status:  http.StatusBadRequest,
		Message: message,
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
	}
}

func (*jsonPresenter) SendUnathorized(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	var response = BaseResponse{
		Status:  http.StatusUnauthorized,
		Message: constant.ErrorUnathorizedMessage,
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
	}
}

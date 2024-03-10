// Package webserver provides the necessary functionality to run a server.
package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/luckysetiawan/curio-qa-api/internal/constant"
)

// jsonPresenter stores presenter logic functions.
type jsonPresenter struct{}

// NewJsonPresenter returns jsonPresenter struct.
func NewJsonPresenter() IPresenterJSON {
	return &jsonPresenter{}
}

// SendSuccess sends a success response with optional data.
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

// SendSuccess sends a success response with optional data and count.
func (*jsonPresenter) SendSuccessWithCount(w http.ResponseWriter, data interface{}, count int) {
	w.Header().Set("Content-Type", "application/json")

	var response = BaseResponse{
		Status:  http.StatusOK,
		Message: constant.SuccessGeneralMessage,
		Data:    data,
		Count:   count,
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
	}
}

// SendError sends an error response with customizable message.
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

// SendUnathorized sends an error unauthorized response.
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

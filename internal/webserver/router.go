package webserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var r = mux.NewRouter()

func NewMuxRouter() *mux.Router {
	return r
}

func SetRouter(method string, url string, handler func(http.ResponseWriter, *http.Request)) {
	r.HandleFunc(fmt.Sprintf("/api%s", url), handler).Methods(method)
}

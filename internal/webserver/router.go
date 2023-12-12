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

func SetRouter(method string, url string, handler func(http.ResponseWriter, *http.Request), accessType []int) {
	if len(accessType) == 0 {
		r.HandleFunc(fmt.Sprintf("/api%s", url), handler).Methods(method)
	} else {
		r.HandleFunc(fmt.Sprintf("/api%s", url), Authenticate(handler, accessType)).Methods(method)
	}
}

// Package webserver provides the necessary functionality to run a server.
package webserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// r contains a mux router.
var r = mux.NewRouter()

// NewRouter returns the created mux router.
func NewMuxRouter() *mux.Router {
	return r
}

// SetRouter adds the url path and handler to the mux router.
func SetRouter(method string, url string, handler func(http.ResponseWriter, *http.Request), accessType []int) {
	if len(accessType) == 0 {
		r.HandleFunc(fmt.Sprintf("/api%s", url), handler).Methods(method)
	} else {
		r.HandleFunc(fmt.Sprintf("/api%s", url), Authenticate(handler, accessType)).Methods(method)
	}
}

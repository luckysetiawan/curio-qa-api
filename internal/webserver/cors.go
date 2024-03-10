// Package webserver provides the necessary functionality to run a server.
package webserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// SetCors returns the http handler with the configured cors.
func SetCors(mux *mux.Router) http.Handler {
	c := cors.Default()

	return c.Handler(mux)
}

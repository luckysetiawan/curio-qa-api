package webserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func SetCors(mux *mux.Router) http.Handler {
	c := cors.Default()

	return c.Handler(mux)
}

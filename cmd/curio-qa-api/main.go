package main

import (
	"net/http"

	_ "github.com/luckysetiawan/curio-qa-api/internal/routes"
	"github.com/luckysetiawan/curio-qa-api/internal/webserver"
)

func main() {
	var handler http.Handler

	router := webserver.NewMuxRouter()
	handler = webserver.SetCors(router)

	webserver.ListenAndServe(":8080", handler)
}

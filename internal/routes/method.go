package routes

import (
	"net/http"

	"github.com/luckysetiawan/curio-qa-api/internal/webserver"
)

func Get(url string, handler func(http.ResponseWriter, *http.Request)) {
	webserver.SetRouter("GET", url, handler)
}

func Put(url string, handler func(http.ResponseWriter, *http.Request)) {
	webserver.SetRouter("PUT", url, handler)
}

func Post(url string, handler func(http.ResponseWriter, *http.Request)) {
	webserver.SetRouter("POST", url, handler)
}

func Delete(url string, handler func(http.ResponseWriter, *http.Request)) {
	webserver.SetRouter("DELETE", url, handler)
}

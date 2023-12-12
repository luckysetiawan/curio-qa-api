package routes

import (
	"net/http"

	"github.com/luckysetiawan/curio-qa-api/internal/webserver"
)

func Get(url string, handler func(http.ResponseWriter, *http.Request), accessType ...int) {
	webserver.SetRouter("GET", url, handler, accessType)
}

func Put(url string, handler func(http.ResponseWriter, *http.Request), accessType ...int) {
	webserver.SetRouter("PUT", url, handler, accessType)
}

func Post(url string, handler func(http.ResponseWriter, *http.Request), accessType ...int) {
	webserver.SetRouter("POST", url, handler, accessType)
}

func Delete(url string, handler func(http.ResponseWriter, *http.Request), accessType ...int) {
	webserver.SetRouter("DELETE", url, handler, accessType)
}

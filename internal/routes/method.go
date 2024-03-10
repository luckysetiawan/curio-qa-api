// Package routes compile the necessary packages and handlers with the defined
// url to be added to the router.
package routes

import (
	"net/http"

	"github.com/luckysetiawan/curio-qa-api/internal/webserver"
)

// Get sets the given url and handler to routes with the GET method.
func Get(url string, handler func(http.ResponseWriter, *http.Request), accessType ...int) {
	webserver.SetRouter("GET", url, handler, accessType)
}

// Post sets the given url and handler to routes with the POST method.
func Post(url string, handler func(http.ResponseWriter, *http.Request), accessType ...int) {
	webserver.SetRouter("POST", url, handler, accessType)
}

// Put sets the given url and handler to routes with the PUT method.
func Put(url string, handler func(http.ResponseWriter, *http.Request), accessType ...int) {
	webserver.SetRouter("PUT", url, handler, accessType)
}

// Delete sets the given url and handler to routes with the DELETE method.
func Delete(url string, handler func(http.ResponseWriter, *http.Request), accessType ...int) {
	webserver.SetRouter("DELETE", url, handler, accessType)
}

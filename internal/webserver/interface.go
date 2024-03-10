// Package webserver provides the necessary functionality to run a server.
package webserver

import "net/http"

// IPresenterJSON defines methods that must be met for a JSON presenter.
type IPresenterJSON interface {
	SendSuccess(w http.ResponseWriter, data ...interface{})
	SendSuccessWithCount(w http.ResponseWriter, data interface{}, count int)
	SendError(w http.ResponseWriter, message string)
	SendUnathorized(w http.ResponseWriter)
}

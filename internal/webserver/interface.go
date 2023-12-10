package webserver

import "net/http"

type IPresenterJSON interface {
	SendSuccess(w http.ResponseWriter, data ...interface{})
	SendError(w http.ResponseWriter, message string)
}

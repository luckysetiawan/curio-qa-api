package webserver

import "net/http"

type IPresenterJSON interface {
	SendSuccess(w http.ResponseWriter, data ...interface{})
	SendSuccessWithCount(w http.ResponseWriter, data interface{}, count int)
	SendError(w http.ResponseWriter, message string)
	SendUnathorized(w http.ResponseWriter)
}

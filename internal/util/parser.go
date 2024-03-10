// Package util provides utility functions to support the server.
package util

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// ParseJSON reads the JSON body then unmarshal the data to the given variables.
func ParseJSON(r *http.Request, v any) error {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}

	return nil
}

// ParsePathParam returns the path parameter sent.
func ParsePathParam(r *http.Request, s string) string {
	vars := mux.Vars(r)

	return vars[s]
}

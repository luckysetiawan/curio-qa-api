package util

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

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

func ParsePathParam(r *http.Request, s string) string {
	vars := mux.Vars(r)

	return vars[s]
}

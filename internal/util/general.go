// Package util provides utility functions to support the server.
package util

import (
	"encoding/json"
	"fmt"
)

// Contains checks whether a value exist in an array of data.
func Contains[T comparable](data []T, value T) bool {
	for _, v := range data {
		if v == value {
			return true
		}
	}

	return false
}

// EncodeJSON marshal the data to string data.
func EncodeJSON(v any) string {
	entity, err := json.Marshal(v)
	if err != nil {
		fmt.Println("error json marshal")
	}

	return string(entity)
}

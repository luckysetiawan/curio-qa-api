package util

import (
	"encoding/json"
	"fmt"
)

func Contains[T comparable](data []T, value T) bool {
	for _, v := range data {
		if v == value {
			return true
		}
	}

	return false
}

func EncodeJSON(v any) string {
	entity, err := json.Marshal(v)
	if err != nil {
		fmt.Println("error json marshal")
	}

	return string(entity)
}

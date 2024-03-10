// Package parser stores all parsing logic the server uses.
package parser

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/luckysetiawan/curio-qa-api/internal/util"
	"github.com/luckysetiawan/curio-qa-api/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestParseUserEntity(t *testing.T) {
	testCases := []struct {
		name          string
		payload       []byte
		expectedData  entity.User
		expectedError bool
	}{
		{
			name: "Success",
			payload: []byte(`{
				"displayName": "User's Display Name",
				"username": "users_username",
				"password": "usersPassword123"
			}`),
			expectedData: entity.User{
				DisplayName: "User's Display Name",
				Username:    "users_username",
				Password:    "usersPassword123",
			},
			expectedError: false,
		},
		{
			name:          "Error",
			payload:       []byte(``),
			expectedData:  entity.User{},
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(tc.payload))
			if err != nil {
				t.Fatal(err)
			}

			parser := NewUserParser()
			actualData, actualError := parser.ParseUserEntity(req)

			assert.JSONEq(t, util.EncodeJSON(tc.expectedData), util.EncodeJSON(actualData))
			if tc.expectedError {
				assert.NotNil(t, actualError)
			} else {
				assert.Nil(t, actualError)
			}
		})
	}
}

func TestParseUsername(t *testing.T) {
	testCases := []struct {
		name     string
		vars     map[string]string
		expected string
	}{
		{
			name:     "Success",
			vars:     map[string]string{"username": "users_username"},
			expected: "users_username",
		},
		{
			name:     "Error",
			vars:     map[string]string{},
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := &http.Request{}
			req = mux.SetURLVars(req, tc.vars)

			parser := NewUserParser()
			actual := parser.ParseUsername(req)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

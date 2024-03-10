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

func TestParseCurioEntity(t *testing.T) {
	testCases := []struct {
		name          string
		payload       []byte
		expectedData  entity.Curio
		expectedError bool
	}{
		{
			name:          "Success",
			payload:       []byte(`{"content": "test"}`),
			expectedData:  entity.Curio{Content: "test"},
			expectedError: false,
		},
		{
			name:          "Error",
			payload:       []byte(``),
			expectedData:  entity.Curio{},
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/curio", bytes.NewBuffer(tc.payload))
			if err != nil {
				t.Fatal(err)
			}

			parser := NewCurioParser()
			actualData, actualError := parser.ParseCurioEntity(req)

			assert.JSONEq(t, util.EncodeJSON(tc.expectedData), util.EncodeJSON(actualData))
			if tc.expectedError {
				assert.NotNil(t, actualError)
			} else {
				assert.Nil(t, actualError)
			}
		})
	}
}

func TestParseCurioID(t *testing.T) {
	testCases := []struct {
		name     string
		vars     map[string]string
		expected string
	}{
		{
			name:     "Success",
			vars:     map[string]string{"curioID": "curioID123"},
			expected: "curioID123",
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

			parser := NewCurioParser()
			actual := parser.ParseCurioID(req)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestParseCurioReceiverUsername(t *testing.T) {
	testCases := []struct {
		name     string
		vars     map[string]string
		expected string
	}{
		{
			name:     "Success",
			vars:     map[string]string{"receiverUsername": "receiver_username"},
			expected: "receiver_username",
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

			parser := NewCurioParser()
			actual := parser.ParseCurioReceiverUsername(req)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

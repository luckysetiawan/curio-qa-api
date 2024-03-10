// Package parser stores all parsing logic the server uses.
package parser

import (
	"net/http"

	"github.com/luckysetiawan/curio-qa-api/internal/util"
	"github.com/luckysetiawan/curio-qa-api/pkg/entity"
)

// curioParser stores parsing logic functions.
type curioParser struct{}

// NewCurioParser returns curioParser struct.
func NewCurioParser() ICurioParser {
	return &curioParser{}
}

// ParseCurioEntity returns curio entity from JSON data.
func (*curioParser) ParseCurioEntity(r *http.Request) (entity.Curio, error) {
	var (
		curio entity.Curio
		err   error
	)
	err = util.ParseJSON(r, &curio)
	if err != nil {
		return entity.Curio{}, err
	}

	return curio, nil
}

// ParseCurioID returns curio ID from path parameter.
func (*curioParser) ParseCurioID(r *http.Request) string {
	ID := util.ParsePathParam(r, "curioID")

	return ID
}

// ParseCurioReceiverUsername returns curio receiver username from path
// parameter.
func (*curioParser) ParseCurioReceiverUsername(r *http.Request) string {
	receiverUsername := util.ParsePathParam(r, "receiverUsername")

	return receiverUsername
}

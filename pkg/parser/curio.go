package parser

import (
	"net/http"

	"github.com/luckysetiawan/curio-qa-api/internal/util"
	"github.com/luckysetiawan/curio-qa-api/pkg/entity"
)

type curioParser struct{}

func NewCurioParser() ICurioParser {
	return &curioParser{}
}

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

func (*curioParser) ParseCurioID(r *http.Request) string {
	ID := util.ParsePathParam(r, "curioID")

	return ID
}

func (*curioParser) ParseCurioReceiverUsername(r *http.Request) string {
	receiverUsername := util.ParsePathParam(r, "receiverUsername")

	return receiverUsername
}

package oracle

import (
	"errors"

	"github.com/fox-one/pando/core"
)

func New() core.OracleService {
	return &oracleService{}
}

type oracleService struct {
}

func (s *oracleService) Parse(b []byte) ([]byte, error) {
	return nil, errors.New("implement me")
}

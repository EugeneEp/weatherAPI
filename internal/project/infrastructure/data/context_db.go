package data

import (
	"github.com/sarulabs/di/v2"
)

func ctxDB(ctn di.Container) (Context, error) {
	return &context{}, nil
}

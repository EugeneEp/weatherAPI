package data

import (
	"github.com/sarulabs/di/v2"
	"projectname/internal/project/infrastructure/data/db/weather"
)

func ctxDB(ctn di.Container) (Context, error) {
	w, err := weather.Context(ctn)

	if err != nil {
		return nil, err
	}

	return &context{
		weather: w,
	}, nil
}

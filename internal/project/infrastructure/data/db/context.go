package db

import (
	"github.com/sarulabs/di/v2"
	"go.uber.org/zap"
	"projectname/internal/project/infrastructure/logger"
)

func Ctx(ctn di.Container) (*Context, error) {
	var log *zap.Logger

	if err := ctn.Fill(logger.BaseServiceName, &log); err != nil {
		return nil, err
	}

	return &Context{
		Log: log,
	}, nil
}

type Context struct {
	Log *zap.Logger
}

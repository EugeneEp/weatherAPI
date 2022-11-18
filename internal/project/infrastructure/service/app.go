package service

import (
	"fmt"
	"github.com/sarulabs/di/v2"
	"go.uber.org/zap"
	"projectname/internal/project/infrastructure/logger"
)

func New(ctn di.Container) (*app, error) {
	var log *zap.Logger

	if err := ctn.Fill(logger.BaseServiceName, &log); err != nil {
		return nil, err
	}

	a := &app{ctn: ctn, errorChannel: make(chan error, 5), logger: log}

	go func() {
		for {
			err := <-a.errorChannel
			log.Error(`Application Error`, zap.Error(err))
		}
	}()

	return a, nil
}

type app struct {
	ctn          di.Container
	logger       *zap.Logger
	errorChannel chan error
}

const name = `project`

func (a *app) Start() error {
	a.logger.Info(fmt.Sprintf(`Starting server %s`, name))

	go a.startWorkers()

	return a.startWeb()
}

package service

import (
	"fmt"
	"github.com/kardianos/service"
	"github.com/sarulabs/di/v2"
	"go.uber.org/zap"
	"os"
	domain "projectname/internal/project/domain/service"
	"projectname/internal/project/infrastructure/logger"
)

func NewApp(ctn di.Container) (service.Service, error) {
	var log *zap.Logger

	if err := ctn.Fill(logger.BaseServiceName, &log); err != nil {
		return nil, err
	}

	a := &app{ctn: ctn, errorChannel: make(chan error, 5), logger: log}

	c := &service.Config{
		Name:        domain.Name,
		DisplayName: domain.DisplayName,
		Description: domain.Description,
	}

	go func() {
		for {
			err := <-a.errorChannel
			log.Error(`Application Error`, zap.Error(err))
		}
	}()

	return service.New(a, c)
}

type app struct {
	ctn          di.Container
	logger       *zap.Logger
	errorChannel chan error
}

const name = `project`

func (a *app) Start(_ service.Service) error {
	a.logger.Info(fmt.Sprintf(`Starting server %s`, name))

	go a.startWorkers()
	go a.startWeb()

	return nil
}

func (a *app) Stop(_ service.Service) error {
	a.logger.Info("Stopping the service " + domain.Name)

	if err := a.ctn.DeleteWithSubContainers(); err != nil {
		a.logger.Error(`Stopping service error`, zap.Error(err))
	}

	a.logger.Info("Dependencies already deleted")

	if service.Interactive() {
		a.logger.Info("app.Stop: os.Exit calling")
		os.Exit(0)
	}

	a.logger.Info("app.Stop: returning nil error")
	return nil
}

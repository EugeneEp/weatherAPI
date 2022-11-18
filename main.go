package main

import (
	"go.uber.org/zap"
	"projectname/internal/project"
	"projectname/internal/project/infrastructure/config"
	"projectname/internal/project/infrastructure/container"
	"projectname/internal/project/infrastructure/logger"
	"projectname/internal/project/infrastructure/service"
)

func main() {
	app := project.New()

	defer func() {
		if r := recover(); r != nil {
			app.Logger().Fatal(`Fatal Application Error`, zap.Reflect(`recovered`, r))
		}
	}()

	cfg := app.Config()

	if err := config.SetDefaults(cfg); err != nil {
		app.Logger().Warn(`Set Defaults Failed`, zap.Error(err))
	}

	if err := config.ReadEnv(cfg); err != nil {
		app.Logger().Warn(`Set Defaults From Env Failed`, zap.Error(err))
	}

	log, err := logger.New(app.Config())

	if err != nil {
		panic(err)
	}

	app.WithLogger(log)

	ctn, err := container.New(app)

	if err != nil {
		app.Logger().Fatal(`Application Error`, zap.Error(err))
		return
	}

	s, err := service.New(ctn)

	if err != nil {
		app.Logger().Fatal(`Application Error`, zap.Error(err))
		return
	}

	if err := s.Start(); err != nil {
		app.Logger().Fatal(`Application Error`, zap.Error(err))
		return
	}
}

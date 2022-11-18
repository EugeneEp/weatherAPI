package container

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di/v2"
	"github.com/spf13/viper"
	"github.com/zhashkevych/scheduler"
	"go.uber.org/zap"
	"projectname/internal/project"
	"projectname/internal/project/infrastructure/background"
	"projectname/internal/project/infrastructure/config"
	"projectname/internal/project/infrastructure/data"
	"projectname/internal/project/infrastructure/database/driver/pgx"
	"projectname/internal/project/infrastructure/logger"
	"projectname/internal/project/interfaces/web/echo/server"
)

// New инициализирует контейнер зависимостей
func New(app project.Project) (di.Container, error) {
	b, err := di.NewBuilder()

	if err != nil {
		return nil, err
	}

	if err = b.Add(initDependencies(app)...); err != nil {
		return nil, err
	}

	return b.Build(), nil
}

// initDependencies инициализирует зависимости проекта
func initDependencies(app project.Project) []di.Def {
	return []di.Def{
		{
			Name: logger.BaseServiceName,
			Build: func(ctn di.Container) (interface{}, error) {
				return app.Logger(), nil
			},
		},
		{
			Name: config.ServiceName,
			Build: func(ctn di.Container) (interface{}, error) {
				return app.Config(), nil
			},
		},
		{
			Name: pgx.ServiceName,
			Build: func(ctn di.Container) (interface{}, error) {
				var cfg *viper.Viper

				if err := ctn.Fill(config.ServiceName, &cfg); err != nil {
					return nil, err
				}

				val, err := pgx.NewDriver(context.Background(), cfg)

				if err != nil {
					logServiceBuildingError(app, pgx.ServiceName, err)
				}

				return val, err
			},
			Close: func(obj interface{}) (err error) {
				logServiceClosing(app, pgx.ServiceName)
				obj.(*pgx.Driver).Close()
				return
			},
		},
		{
			Name: data.ServiceName,
			Build: func(ctn di.Container) (interface{}, error) {
				val, err := data.New(ctn)

				if err != nil {
					logServiceBuildingError(app, data.ServiceName, err)
				}

				return val, err
			},
		},
		{
			Name: server.ServiceName,
			Build: func(ctn di.Container) (interface{}, error) {
				var log *zap.Logger

				if err := ctn.Fill(logger.BaseServiceName, &log); err != nil {
					return nil, err
				}

				val, err := server.New(ctn, log)

				if err != nil {
					logServiceBuildingError(app, server.ServiceName, err)
				}

				return val, err
			},
			Close: func(obj interface{}) (err error) {
				logServiceClosing(app, server.ServiceName)

				if err = obj.(*echo.Echo).Close(); err != nil {
					logServiceClosingError(app, server.ServiceName, err)
				}

				return
			},
		},
		{
			Name: background.ServiceName,
			Build: func(ctn di.Container) (interface{}, error) {
				var cfg *viper.Viper

				if err := ctn.Fill(config.ServiceName, &cfg); err != nil {
					return nil, err
				}

				val, err := background.New(ctn, cfg)

				if err != nil {
					logServiceBuildingError(app, background.ServiceName, err)
				}

				return val, err
			},
			Close: func(obj interface{}) (err error) {
				logServiceClosing(app, background.ServiceName)
				obj.(*scheduler.Scheduler).Stop()
				return
			},
		},
	}
}

func logServiceBuildingError(app project.Project, srv string, err error) {
	app.Logger().Error(`Service Building Error`, zap.String(`service_name`, srv), zap.Error(err))
}

func logServiceClosing(app project.Project, srv string) {
	app.Logger().Info(`Service Closing Message`, zap.String(`service_name`, srv))
}

func logServiceClosingError(app project.Project, srv string, err error) {
	app.Logger().Error(`Service Closing Error`, zap.String(`service_name`, srv), zap.Error(err))
}

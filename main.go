package main

import (
	"go.uber.org/zap"
	"projectname/internal/project"
	"projectname/internal/project/infrastructure/logger"
	"projectname/internal/project/interfaces/cli"
	"projectname/internal/project/interfaces/cli/cmd"
)

func main() {
	app := project.New()

	log, err := logger.New(app.Config())

	if err != nil {
		panic(err)
	}

	if r := recover(); r != nil {
		app.Logger().Fatal(`Fatal Application Error`, zap.Reflect(`recovered`, r))
		return
	}

	app.WithLogger(log)

	root := cmd.Run(app)

	cli.Bind(root, app)

	if err := root.Execute(); err != nil {
		app.Logger().Fatal(err.Error())
		return
	}
}

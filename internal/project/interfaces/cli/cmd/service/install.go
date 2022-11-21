package service

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"projectname/internal/project"
	"projectname/internal/project/infrastructure/config"
	"projectname/internal/project/infrastructure/container"
	"projectname/internal/project/infrastructure/logger"
	"projectname/internal/project/infrastructure/service"
	"projectname/internal/project/interfaces/cli/ident"
)

func Install(app project.Project) *cobra.Command {
	cmd := &cobra.Command{
		Use: ident.CmdInstall,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			cfg := app.Config()

			if err := config.SetDefaults(cfg); err != nil {
				app.Logger().Warn(ident.MsgSetDefaultsFailed, zap.Error(err))
			}

			app.WithConfig(cfg)

			log, err := logger.New(cfg)

			if err != nil {
				app.Logger().Warn(ident.MsgInitializeAppLoggerError, zap.Error(err))
			} else {
				app.WithLogger(log)
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctn, err := container.New(app)

			if err != nil {
				return err
			}

			s, err := service.NewApp(ctn)

			if err != nil {
				app.Logger().Fatal(err.Error())
				return err
			}

			if err = s.Install(); err != nil {
				app.Logger().Error(err.Error())
				return err
			}

			return err
		},
	}

	return cmd
}

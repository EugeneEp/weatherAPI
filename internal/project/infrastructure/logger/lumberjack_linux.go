package logger

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/url"
	"projectname/internal/project/domain/configuration"
)

type lumberjackSink struct {
	*lumberjack.Logger
}

func (lumberjackSink) Sync() error { return nil }

func registerLumberjackSink(cfg *viper.Viper) error {
	return zap.RegisterSink("lumberjack", func(u *url.URL) (zap.Sink, error) {
		return lumberjackSink{
			Logger: &lumberjack.Logger{
				Filename:   u.Path,
				MaxSize:    cfg.GetInt(configuration.LogMaxSize),
				MaxBackups: cfg.GetInt(configuration.LogMaxBackups),
				MaxAge:     cfg.GetInt(configuration.LogMaxAge),
			},
		}, nil
	})
}

package logger

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"path/filepath"
	"projectname/internal/project/domain/configuration"
)

const BaseServiceName = `DefaultAppLogger`

// New Инициализирует сервис логирования, в зависимости от настроек конфигурации
func New(cfg *viper.Viper) (*zap.Logger, error) {
	logLevel := cfg.GetString(configuration.LogLevel)
	logDir := cfg.GetString(configuration.DirLog)
	logName := cfg.GetString(configuration.FileLogName)

	var level zap.AtomicLevel

	switch logLevel {
	case `debug`:
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case `warning`:
		level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case `error`:
		level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case `panic`:
		level = zap.NewAtomicLevelAt(zap.PanicLevel)
	case `fatal`:
		level = zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	developmentMode := cfg.GetBool(configuration.DevelopmentMode)

	if logName == "" || logDir == "" {
		if developmentMode {
			return zap.NewDevelopment(zap.IncreaseLevel(level))
		} else {
			return zap.NewProduction(zap.IncreaseLevel(level))
		}
	}

	var conf zap.Config

	if err := registerLumberjackSink(cfg); err != nil {
		return nil, err
	}

	path := "lumberjack:" + filepath.Join(logDir, logName)

	if developmentMode {
		conf = zap.NewDevelopmentConfig()
	} else {
		conf = zap.NewProductionConfig()
	}

	conf.Level = level
	conf.OutputPaths = []string{path}
	conf.ErrorOutputPaths = []string{path}

	return conf.Build()
}

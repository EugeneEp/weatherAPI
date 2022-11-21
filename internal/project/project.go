package project

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"projectname/internal/project/infrastructure/config"
)

type (
	Project interface {
		Logger() *zap.Logger
		Config() *viper.Viper

		WithLogger(log *zap.Logger)
		WithConfig(cfg *viper.Viper)
	}

	project struct {
		log *zap.Logger
		cfg *viper.Viper
	}
)

func New() Project {
	return &project{
		log: zap.L(),
		cfg: config.New(),
	}
}

func (c project) Logger() *zap.Logger { return c.log }

func (c project) Config() *viper.Viper { return c.cfg }

func (c *project) WithLogger(log *zap.Logger) {
	c.log = log
}

func (c *project) WithConfig(cfg *viper.Viper) {
	c.cfg = cfg
}

package data

import (
	"github.com/sarulabs/di/v2"
	"github.com/spf13/viper"
	"projectname/internal/project/infrastructure/config"
)

type (
	Context interface {
	}

	context struct {
	}
)

const ServiceName = `DataStorageService`

// New Инициализирует сервис хранения данных, в зависимости от настроек конфигурации
func New(ctn di.Container) (Context, error) {
	var cfg *viper.Viper

	if err := ctn.Fill(config.ServiceName, &cfg); err != nil {
		return nil, err
	}

	return ctxDB(ctn)
}

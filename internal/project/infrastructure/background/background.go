package background

import (
	"github.com/sarulabs/di/v2"
	"github.com/spf13/viper"
	"github.com/zhashkevych/scheduler"
)

const ServiceName = `BackgroundServiceName`

func New(ctn di.Container, cfg *viper.Viper) (*scheduler.Scheduler, error) {
	background := scheduler.NewScheduler()

	return background, nil
}

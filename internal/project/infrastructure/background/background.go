package background

import (
	"context"
	"github.com/sarulabs/di/v2"
	"github.com/spf13/viper"
	"github.com/zhashkevych/scheduler"
	"projectname/internal/project/core/weather"
	"projectname/internal/project/domain/configuration"
	"time"
)

const ServiceName = `BackgroundServiceName`

func New(ctn di.Container, cfg *viper.Viper) (*scheduler.Scheduler, error) {
	getCityWeatherTime := time.Duration(cfg.GetInt(configuration.SyncGetCityWeatherTime))
	cntDayArchiveTime := time.Duration(cfg.GetInt(configuration.SyncСntDayArchiveTime))

	background := scheduler.NewScheduler()
	background.Add(context.Background(), func(ctx context.Context) {
		weather.Write(ctn)
	}, time.Minute*getCityWeatherTime)

	background.Add(context.Background(), func(ctx context.Context) {
		weather.SyncGetAvgTemperature(ctn)
	}, 24*time.Hour*cntDayArchiveTime)

	return background, nil
}

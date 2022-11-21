package weather

import (
	"fmt"
	"github.com/sarulabs/di/v2"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"projectname/internal/project/domain/configuration"
	domain "projectname/internal/project/domain/open_weather"
	"projectname/internal/project/domain/weather"
	"projectname/internal/project/infrastructure/alerts"
	"projectname/internal/project/infrastructure/config"
	"projectname/internal/project/infrastructure/data"
	"projectname/internal/project/infrastructure/logger"
	openWeather "projectname/internal/project/infrastructure/open_weather"
	"time"
)

func Create(ctn di.Container, rq weather.Request) error {
	var (
		api *openWeather.OpenWeatherApi
		ctx data.Context
	)

	if err := ctn.Fill(openWeather.ServiceName, &api); err != nil {
		return err
	}

	city, err := api.GetCityInfo(rq.City)
	if err != nil {
		return err
	}

	if err := ctn.Fill(data.ServiceName, &ctx); err != nil {
		return err
	}

	return ctx.Weather().Create(*city)
}

func Delete(ctn di.Container, rq weather.Request) error {
	var (
		ctx data.Context
	)

	if err := ctn.Fill(data.ServiceName, &ctx); err != nil {
		return err
	}

	return ctx.Weather().Delete(rq)
}

func Get(ctn di.Container) (*weather.Get, error) {
	var (
		ctx data.Context
	)

	if err := ctn.Fill(data.ServiceName, &ctx); err != nil {
		return nil, err
	}

	return ctx.Weather().Get()
}

func Write(ctn di.Container) {
	var (
		api *openWeather.OpenWeatherApi
		ctx data.Context
		log *zap.Logger
		al  *alerts.Alerts
	)

	if err := ctn.Fill(logger.BaseServiceName, &log); err != nil {
		return
	}

	if err := ctn.Fill(openWeather.ServiceName, &api); err != nil {
		log.Warn(err.Error(), zap.Error(err))
		return
	}

	if err := ctn.Fill(data.ServiceName, &ctx); err != nil {
		log.Warn(err.Error(), zap.Error(err))
		return
	}

	if err := ctn.Fill(alerts.ServiceName, &al); err != nil {
		log.Warn(err.Error(), zap.Error(err))
		return
	}

	c, err := Get(ctn)
	if err != nil {
		log.Warn(err.Error(), zap.Error(err))
		return
	}

	for _, v := range c.Cities {
		go func(city weather.City) {
			temp, err := api.GetWeather(domain.City{
				Name: city.Name,
				Lat:  city.Lat,
				Lon:  city.Lon,
			})
			if err != nil {
				log.Warn(`Get `+city.Name+`Temp Failed`, zap.Error(err))
			}

			al.Alert(`City: ` + city.Name + ` | Current temp:` + fmt.Sprintf("%.6f", *temp))

			if err := ctx.Weather().Write(weather.Temperature{
				CityName: city.Name,
				Temp:     *temp,
				Dt:       time.Now().Unix(),
			}); err != nil {
				log.Warn(`Get `+city.Name+`Temp Failed`, zap.Error(err))
			}
		}(v)
	}
}

func GetAvgTemperature(ctn di.Container, req weather.Request) (*weather.AvgTemp, error) {
	var (
		ctx data.Context
		cfg *viper.Viper
	)

	if err := ctn.Fill(data.ServiceName, &ctx); err != nil {
		return nil, err
	}

	if err := ctn.Fill(config.ServiceName, &cfg); err != nil {
		return nil, err
	}

	endDate := time.Now().Unix()
	cntDay := cfg.GetInt(configuration.СntDayArchive)
	avgTemp := weather.AvgTemp{
		City:      req.City,
		CntDay:    cntDay,
		StartDate: endDate - int64(cntDay*24*60*60),
		EndDate:   endDate,
	}
	temp, err := ctx.Weather().GetAvgTemp(avgTemp)
	if err != nil {
		return nil, err
	}

	avgTemp.Temp = *temp

	return &avgTemp, nil
}

func SyncGetAvgTemperature(ctn di.Container) {
	var (
		ctx data.Context
		cfg *viper.Viper
		log *zap.Logger
	)

	if err := ctn.Fill(logger.BaseServiceName, &log); err != nil {
		return
	}

	if err := ctn.Fill(data.ServiceName, &ctx); err != nil {
		log.Warn(`Get Avg Temp Failed`, zap.Error(err))
		return
	}

	if err := ctn.Fill(config.ServiceName, &cfg); err != nil {
		log.Warn(`Get Avg Temp Failed`, zap.Error(err))
		return
	}

	c, err := Get(ctn)
	if err != nil {
		log.Warn(err.Error(), zap.Error(err))
		return
	}

	for _, v := range c.Cities {
		go func(city weather.City) {
			endDate := time.Now().Unix()
			cntDay := cfg.GetInt(configuration.СntDayArchive)
			avgTemp := weather.AvgTemp{
				City:      city.Name,
				CntDay:    cntDay,
				StartDate: endDate - int64(cntDay*24*60*60),
				EndDate:   endDate,
			}
			temp, err := ctx.Weather().GetAvgTemp(avgTemp)
			if err != nil {
				log.Warn(`Get Avg Temp Failed`, zap.Error(err))
				return
			}

			avgTemp.Temp = *temp

			if err := ctx.Weather().WriteAvg(avgTemp); err != nil {
				log.Warn(`Get Avg Temp Failed`, zap.Error(err))
				return
			}
		}(v)
	}
}

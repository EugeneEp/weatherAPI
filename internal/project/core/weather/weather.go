package weather

import (
	"github.com/sarulabs/di/v2"
	"go.uber.org/zap"
	domain "projectname/internal/project/domain/open_weather"
	"projectname/internal/project/domain/weather"
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
	)

	if err := ctn.Fill(logger.BaseServiceName, &log); err != nil {
		return
	}

	if err := ctn.Fill(openWeather.ServiceName, &api); err != nil {
		log.Warn(err.Error(), zap.Error(err))
		return
	}

	c, err := Get(ctn)
	if err != nil {
		log.Warn(err.Error(), zap.Error(err))
		return
	}

	if err := ctn.Fill(data.ServiceName, &ctx); err != nil {
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

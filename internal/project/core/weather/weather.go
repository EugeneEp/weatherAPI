package weather

import (
	"github.com/sarulabs/di/v2"
	"projectname/internal/project/domain/weather"
	"projectname/internal/project/infrastructure/data"
	openWeather "projectname/internal/project/infrastructure/open_weather"
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

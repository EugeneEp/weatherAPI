package open_weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"projectname/internal/project/domain/configuration"
	domain "projectname/internal/project/domain/open_weather"
)

const (
	ServiceName     = `OpenWeatherApi`
	cityInfoPath    = `geo/1.0/direct`
	weatherInfoPath = `data/2.5/weather`
)

type OpenWeatherApi struct {
	cfg *viper.Viper
	log *zap.Logger
}

func New(cfg *viper.Viper) *OpenWeatherApi {
	return &OpenWeatherApi{
		cfg: cfg,
	}
}

func (o *OpenWeatherApi) GetCityInfo(city string) (*domain.City, error) {
	r, err := http.NewRequest("GET", o.cfg.GetString(configuration.OpenWeatherUrl)+cityInfoPath, nil)
	if err != nil {
		return nil, err
	}

	q := r.URL.Query()
	q.Add("q", city)
	q.Add("limit", "1")
	q.Add("appid", o.cfg.GetString(configuration.OpenWeatherKey))
	r.URL.RawQuery = q.Encode()

	res, err := http.Get(r.URL.String())
	if err != nil {
		return nil, err
	}

	var cities *[]domain.City

	if err = json.NewDecoder(res.Body).Decode(&cities); err != nil {
		return nil, err
	}

	if cities != nil && len(*cities) > 0 {
		c := *cities
		return &c[0], nil
	}

	return nil, errors.New(`app.OpenWeatherApi: City Not Found`)
}

func (o *OpenWeatherApi) GetWeather(city domain.City) (*float64, error) {
	r, err := http.NewRequest("GET", o.cfg.GetString(configuration.OpenWeatherUrl)+weatherInfoPath, nil)
	if err != nil {
		return nil, err
	}

	q := r.URL.Query()
	q.Add("lat", fmt.Sprintf("%.6f", city.Lat))
	q.Add("lon", fmt.Sprintf("%.6f", city.Lon))
	q.Add("units", "metric")
	q.Add("appid", o.cfg.GetString(configuration.OpenWeatherKey))
	r.URL.RawQuery = q.Encode()

	res, err := http.Get(r.URL.String())
	if err != nil {
		return nil, err
	}

	var weather *domain.Weather

	if err = json.NewDecoder(res.Body).Decode(&weather); err != nil {
		return nil, err
	}

	if weather != nil {
		return &weather.Main.Temp, nil
	}

	return nil, errors.New(`app.OpenWeatherApi: Weather Not Found`)
}

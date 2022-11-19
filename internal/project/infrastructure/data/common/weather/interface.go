package weather

import (
	"projectname/internal/project/domain/open_weather"
	"projectname/internal/project/domain/weather"
)

type Interface interface {
	Create(rq open_weather.City) error
	Delete(rq weather.Request) error
	Write(rq weather.Temperature) error
	Get() (*weather.Get, error)
}

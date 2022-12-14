package weather

import (
	"projectname/internal/project/domain/open_weather"
	"projectname/internal/project/domain/weather"
	"projectname/internal/project/infrastructure/database/driver"
)

type Interface interface {
	Create(rq open_weather.City) error
	Delete(rq weather.Request) error
	Write(rq weather.Temperature) error
	Get() (*weather.Get, error)
	GetAvgTemp(req weather.AvgTemp) (*float64, error)
	WriteAvg(rq weather.AvgTemp) error
	SetDriver(d driver.Driver)
}

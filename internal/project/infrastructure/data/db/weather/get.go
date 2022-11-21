package weather

import (
	"go.uber.org/zap"
	"projectname/internal/project/domain/weather"
)

const queryGetValue = `SELECT name, lat, lon FROM "weather".cities`

func (c *context) Get() (*weather.Get, error) {
	var res []weather.City

	if err := c.driver.Select(&res,
		queryGetValue,
	); err != nil {
		c.reg.Log.Error(weather.ErrNotFound.Error(), zap.Error(err))
		return nil, weather.ErrNotFound
	}

	return &weather.Get{
		Count:  len(res),
		Cities: res,
	}, nil
}

const queryGetAvgValue = `SELECT AVG(temp) as temp FROM "weather".temperature WHERE city_name = $1 AND dt > $2`

func (c *context) GetAvgTemp(req weather.AvgTemp) (*float64, error) {
	var res map[string]float64

	if err := c.driver.Get(&res,
		queryGetAvgValue,
		req.City,
		req.StartDate,
	); err != nil {
		c.reg.Log.Error(weather.ErrTempNotFound.Error(), zap.Error(err))
		return nil, weather.ErrTempNotFound
	}

	if v, ok := res["temp"]; ok {
		return &v, nil
	}

	return nil, weather.ErrTempNotFound
}

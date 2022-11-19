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

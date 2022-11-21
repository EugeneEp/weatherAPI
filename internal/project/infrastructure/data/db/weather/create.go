package weather

import (
	"go.uber.org/zap"
	"projectname/internal/project/domain/open_weather"
	domain "projectname/internal/project/domain/weather"
)

const queryAddValue = `INSERT INTO "weather".cities (name, lat, lon) VALUES ($1, $2, $3)`

func (c *context) Create(rq open_weather.City) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if err := c.driver.Query(
		queryAddValue,
		rq.Name,
		rq.Lat,
		rq.Lon,
	); err != nil {
		c.reg.Log.Error(domain.ErrNotCreated.Error(), zap.Error(err))
		return domain.ErrNotCreated
	}

	return nil
}

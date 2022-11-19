package weather

import (
	"go.uber.org/zap"
	domain "projectname/internal/project/domain/weather"
)

const queryWriteValue = `INSERT INTO "weather".temperature (temp, dt, city_name) VALUES ($1, $2, $3)`

func (c *context) Write(rq domain.Temperature) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if err := c.driver.Query(
		queryWriteValue,
		rq.Temp,
		rq.Dt,
		rq.CityName,
	); err != nil {
		c.reg.Log.Error(domain.ErrNotWritten.Error(), zap.Error(err))
		return domain.ErrNotWritten
	}

	return nil
}

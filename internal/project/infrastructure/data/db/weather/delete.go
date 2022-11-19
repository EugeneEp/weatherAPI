package weather

import (
	"go.uber.org/zap"
	"projectname/internal/project/domain/weather"
)

const queryDelete = `DELETE FROM "weather".cities WHERE name = $1`

func (c *context) Delete(rq weather.Request) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if err := c.driver.Query(
		queryDelete,
		rq.City,
	); err != nil {
		c.reg.Log.Error(weather.ErrNotDeleted.Error(), zap.Error(err))
		return weather.ErrNotDeleted
	}

	return nil
}

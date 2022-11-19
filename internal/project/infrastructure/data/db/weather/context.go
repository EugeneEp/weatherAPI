package weather

import (
	"github.com/sarulabs/di/v2"
	"projectname/internal/project/infrastructure/data/common/weather"
	"projectname/internal/project/infrastructure/data/db"
	"projectname/internal/project/infrastructure/database/driver"
	"projectname/internal/project/infrastructure/database/driver/pgx"
	"sync"
)

type (
	context struct {
		ctn    di.Container
		reg    *db.Context
		mutex  sync.RWMutex
		driver driver.Driver
	}
)

func Context(ctn di.Container) (weather.Interface, error) {
	reg, err := db.Ctx(ctn)

	if err != nil {
		return nil, err
	}

	var d *pgx.Driver

	if err = ctn.Fill(pgx.ServiceName, &d); err != nil {
		return nil, err
	}

	return &context{
		ctn:    ctn,
		reg:    reg,
		driver: d,
	}, nil
}

package database

import (
	"github.com/bxcodec/faker"
	"github.com/sarulabs/di/v2"
	stdLog "log"
	"projectname/internal/project"
	"projectname/internal/project/domain/weather"
	"projectname/internal/project/infrastructure/config"
	"projectname/internal/project/infrastructure/database"
	"projectname/internal/project/infrastructure/database/driver/pgx"
	"projectname/internal/project/infrastructure/logger"
	"projectname/pkg/test"
	"testing"
)

const (
	getCitiesSuccess = `GetCities Success`
	getCitiesErr     = `GetCities Err`
	getCitiesDbErr   = `GetCities Database Error`
)

type (
	getCitiesSuccessPgxDriver struct{ pgx.Driver }
	getCitiesErrPgxDriver     struct{ pgx.Driver }
	getCitiesDbErrPgxDriver   struct{ pgx.Driver }
)

func (d getCitiesSuccessPgxDriver) Select(dst interface{}, _ string, _ ...interface{}) error {
	var p = dst.(*[]weather.City)
	*p = append(*p, gt)
	return nil
}

func (d getCitiesErrPgxDriver) Select(dst interface{}, _ string, _ ...interface{}) error {
	return weather.ErrNotFound
}

func (d getCitiesDbErrPgxDriver) Select(_ interface{}, _ string, _ ...interface{}) error {
	return weather.ErrNotFound
}

var (
	gt weather.City
)

func init() {
	if err := faker.FakeData(&gt); err != nil {
		stdLog.Fatalf("Filling test data error: %s", err)
	}
}

func TestGetCities(t *testing.T) {
	var (
		success = &weather.Get{
			Count: 1,
			Cities: []weather.City{
				gt,
			},
		}

		specificNil = (*weather.Get)(nil)

		cases = []database.TestCase{
			{
				Driver:  getCitiesSuccessPgxDriver{},
				Request: nil,
				Iface:   test.New(nil, getCitiesSuccess, success),
			},

			{
				Driver:  getCitiesErrPgxDriver{},
				Request: nil,
				Iface:   test.New(weather.ErrNotFound, getCitiesErr, specificNil),
			},

			{
				Driver:  getCitiesDbErrPgxDriver{},
				Request: nil,
				Iface:   test.New(weather.ErrNotFound, getCitiesDbErr, specificNil),
			},
		}
	)
	for _, c := range cases {
		ctx.Weather().SetDriver(c.Driver)
		result, err := ctx.Weather().Get()
		c.Iface.Compare(t, result, err)
	}
}

func newCtn(app project.Project) (di.Container, error) {
	b, err := di.NewBuilder()

	if err != nil {
		return nil, err
	}

	if err = b.Add(initDependencies(app)...); err != nil {
		return nil, err
	}

	return b.Build(), nil
}

func initDependencies(app project.Project) []di.Def {
	return []di.Def{
		{
			Name: logger.BaseServiceName,
			Build: func(ctn di.Container) (interface{}, error) {
				return app.Logger(), nil
			},
		},
		{
			Name: config.ServiceName,
			Build: func(ctn di.Container) (interface{}, error) {
				return app.Config(), nil
			},
		},
		{
			Name: pgx.ServiceName,
			Build: func(ctn di.Container) (interface{}, error) {
				return &pgx.Driver{}, nil
			},
		},
	}
}

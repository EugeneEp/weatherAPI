package database

import (
	"projectname/internal/project/domain/open_weather"
	"projectname/internal/project/domain/weather"
	"projectname/internal/project/infrastructure/database"
	"projectname/internal/project/infrastructure/database/driver/pgx"
	"projectname/pkg/test"
	"testing"
)

const (
	createCitySuccess = `CreateCity Success`
	createCityErr     = `CreateCity Err`
	createCityDbErr   = `CreateCity Database Error`
)

type (
	createCitySuccessPgxDriver struct{ pgx.Driver }
	createCityErrPgxDriver     struct{ pgx.Driver }
	createCityDbErrPgxDriver   struct{ pgx.Driver }
)

func (d createCitySuccessPgxDriver) Query(_ string, _ ...interface{}) error {
	return nil
}

func (d createCityErrPgxDriver) Query(_ string, _ ...interface{}) error {
	return weather.ErrNotCreated
}

func (d createCityDbErrPgxDriver) Query(_ string, _ ...interface{}) error {
	return weather.ErrNotCreated
}

func TestCreateCity(t *testing.T) {
	var (
		req = open_weather.City{
			Name: "London",
			Lat:  123,
			Lon:  123,
		}

		cases = []database.TestCase{
			{
				Driver:  createCitySuccessPgxDriver{},
				Request: req,
				Iface:   test.New(nil, createCitySuccess, nil),
			},

			{
				Driver: createCityErrPgxDriver{},
				Request: open_weather.City{
					Name: "waddadwadw",
					Lat:  12312312,
					Lon:  123122,
				},
				Iface: test.New(weather.ErrNotCreated, createCityErr, nil),
			},

			{
				Driver: createCityDbErrPgxDriver{},
				Request: open_weather.City{
					Name: "waddadwadw",
					Lat:  12312312,
					Lon:  123122,
				},
				Iface: test.New(weather.ErrNotCreated, createCityDbErr, nil),
			},
		}
	)
	for _, c := range cases {
		ctx.Weather().SetDriver(c.Driver)
		r := c.Request.(open_weather.City)
		err := ctx.Weather().Create(r)
		c.Iface.Compare(t, nil, err)
	}
}

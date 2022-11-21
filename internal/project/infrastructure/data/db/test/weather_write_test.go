package database

import (
	"projectname/internal/project/domain/weather"
	"projectname/internal/project/infrastructure/database"
	"projectname/internal/project/infrastructure/database/driver/pgx"
	"projectname/pkg/test"
	"testing"
)

const (
	writeTemperatureSuccess = `WriteTemperature Success`
	writeTemperatureErr     = `WriteTemperature Err`
	writeTemperatureDbErr   = `WriteTemperature Database Error`
)

type (
	writeTemperatureSuccessPgxDriver struct{ pgx.Driver }
	writeTemperatureErrPgxDriver     struct{ pgx.Driver }
	writeTemperatureDbErrPgxDriver   struct{ pgx.Driver }
)

func (d writeTemperatureSuccessPgxDriver) Query(_ string, _ ...interface{}) error {
	return nil
}

func (d writeTemperatureErrPgxDriver) Query(_ string, _ ...interface{}) error {
	return weather.ErrNotWritten
}

func (d writeTemperatureDbErrPgxDriver) Query(_ string, _ ...interface{}) error {
	return weather.ErrNotWritten
}

func TestWriteTemperature(t *testing.T) {
	var (
		req = weather.Temperature{
			CityName: "London",
			Temp:     23.2,
			Dt:       124124211,
		}

		cases = []database.TestCase{
			{
				Driver:  writeTemperatureSuccessPgxDriver{},
				Request: req,
				Iface:   test.New(nil, writeTemperatureSuccess, nil),
			},

			{
				Driver: writeTemperatureErrPgxDriver{},
				Request: weather.Temperature{
					CityName: "wadwadwadw",
					Temp:     55.2,
					Dt:       242,
				},
				Iface: test.New(weather.ErrNotWritten, writeTemperatureErr, nil),
			},

			{
				Driver: writeTemperatureDbErrPgxDriver{},
				Request: weather.Temperature{
					CityName: "wadwadwadw",
					Temp:     55.2,
					Dt:       242,
				},
				Iface: test.New(weather.ErrNotWritten, writeTemperatureDbErr, nil),
			},
		}
	)
	for _, c := range cases {
		ctx.Weather().SetDriver(c.Driver)
		r := c.Request.(weather.Temperature)
		err := ctx.Weather().Write(r)
		c.Iface.Compare(t, nil, err)
	}
}

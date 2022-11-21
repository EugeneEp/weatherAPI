package database

import (
	"projectname/internal/project/domain/weather"
	"projectname/internal/project/infrastructure/database"
	"projectname/internal/project/infrastructure/database/driver/pgx"
	"projectname/pkg/test"
	"testing"
)

const (
	writeAvgTemperatureSuccess = `WriteAvgTemperature Success`
	writeAvgTemperatureErr     = `WriteAvgTemperature Err`
	writeAvgTemperatureDbErr   = `WriteAvgTemperature Database Error`
)

type (
	writeAvgTemperatureSuccessPgxDriver struct{ pgx.Driver }
	writeAvgTemperatureErrPgxDriver     struct{ pgx.Driver }
	writeAvgTemperatureDbErrPgxDriver   struct{ pgx.Driver }
)

func (d writeAvgTemperatureSuccessPgxDriver) Query(_ string, _ ...interface{}) error {
	return nil
}

func (d writeAvgTemperatureErrPgxDriver) Query(_ string, _ ...interface{}) error {
	return weather.ErrNotWritten
}

func (d writeAvgTemperatureDbErrPgxDriver) Query(_ string, _ ...interface{}) error {
	return weather.ErrNotWritten
}

func TestWriteAvgTemperature(t *testing.T) {
	var (
		req = weather.AvgTemp{
			City:      "London",
			CntDay:    1,
			StartDate: 123123213,
			EndDate:   142212122,
			Temp:      23.1,
		}

		cases = []database.TestCase{
			{
				Driver:  writeAvgTemperatureSuccessPgxDriver{},
				Request: req,
				Iface:   test.New(nil, writeAvgTemperatureSuccess, nil),
			},

			{
				Driver: writeAvgTemperatureErrPgxDriver{},
				Request: weather.AvgTemp{
					City:      "awdaaddw",
					CntDay:    1,
					StartDate: 124,
					EndDate:   123,
					Temp:      11.1,
				},
				Iface: test.New(weather.ErrNotWritten, writeAvgTemperatureErr, nil),
			},

			{
				Driver: writeAvgTemperatureDbErrPgxDriver{},
				Request: weather.AvgTemp{
					City:      "awdaaddw",
					CntDay:    1,
					StartDate: 124,
					EndDate:   123,
					Temp:      11.1,
				},
				Iface: test.New(weather.ErrNotWritten, writeAvgTemperatureDbErr, nil),
			},
		}
	)
	for _, c := range cases {
		ctx.Weather().SetDriver(c.Driver)
		r := c.Request.(weather.AvgTemp)
		err := ctx.Weather().WriteAvg(r)
		c.Iface.Compare(t, nil, err)
	}
}

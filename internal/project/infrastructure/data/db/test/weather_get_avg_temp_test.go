package database

import (
	"projectname/internal/project/domain/weather"
	"projectname/internal/project/infrastructure/database"
	"projectname/internal/project/infrastructure/database/driver/pgx"
	"projectname/pkg/test"
	"testing"
)

const (
	getAvgTempSuccess = `GetAvgTemp Success`
	getAvgTempErr     = `GetAvgTemp Err`
	getAvgTempDbErr   = `GetAvgTemp Database Error`
)

type (
	getAvgTempSuccessPgxDriver struct{ pgx.Driver }
	getAvgTempErrPgxDriver     struct{ pgx.Driver }
	getAvgTempDbErrPgxDriver   struct{ pgx.Driver }
)

func (d getAvgTempSuccessPgxDriver) Get(dst interface{}, _ string, _ ...interface{}) error {
	var p = dst.(*map[string]float64)
	*p = av
	return nil
}

func (d getAvgTempErrPgxDriver) Get(dst interface{}, _ string, _ ...interface{}) error {
	return weather.ErrTempNotFound
}

func (d getAvgTempDbErrPgxDriver) Get(_ interface{}, _ string, _ ...interface{}) error {
	return weather.ErrTempNotFound
}

var (
	av = make(map[string]float64, 0)
)

func init() {
	av["temp"] = 3.12
}

func TestGetAvgTemp(t *testing.T) {
	var (
		success = 3.12

		specificNil = (*float64)(nil)

		req = weather.AvgTemp{
			City:      "London",
			StartDate: 1232312,
		}

		cases = []database.TestCase{
			{
				Driver:  getAvgTempSuccessPgxDriver{},
				Request: req,
				Iface:   test.New(nil, getAvgTempSuccess, &success),
			},

			{
				Driver: getAvgTempErrPgxDriver{},
				Request: weather.AvgTemp{
					City:      "wdwadwadw",
					StartDate: 1232312,
				},
				Iface: test.New(weather.ErrTempNotFound, getAvgTempErr, specificNil),
			},

			{
				Driver: getAvgTempDbErrPgxDriver{},
				Request: weather.AvgTemp{
					City:      "wdwadwadw",
					StartDate: 1232312,
				},
				Iface: test.New(weather.ErrTempNotFound, getAvgTempDbErr, specificNil),
			},
		}
	)
	for _, c := range cases {
		ctx.Weather().SetDriver(c.Driver)
		r := c.Request.(weather.AvgTemp)
		result, err := ctx.Weather().GetAvgTemp(r)
		c.Iface.Compare(t, result, err)
	}
}

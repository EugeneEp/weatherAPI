package database

import (
	"projectname/internal/project/domain/weather"
	"projectname/internal/project/infrastructure/database"
	"projectname/internal/project/infrastructure/database/driver/pgx"
	"projectname/pkg/test"
	"testing"
)

const (
	deleteCitySuccess = `DeleteCity Success`
	deleteCityErr     = `DeleteCity Err`
	deleteCityDbErr   = `DeleteCity Database Error`
)

type (
	deleteCitySuccessPgxDriver struct{ pgx.Driver }
	deleteCityErrPgxDriver     struct{ pgx.Driver }
	deleteCityDbErrPgxDriver   struct{ pgx.Driver }
)

func (d deleteCitySuccessPgxDriver) Query(_ string, _ ...interface{}) error {
	return nil
}

func (d deleteCityErrPgxDriver) Query(_ string, _ ...interface{}) error {
	return weather.ErrNotDeleted
}

func (d deleteCityDbErrPgxDriver) Query(_ string, _ ...interface{}) error {
	return weather.ErrNotDeleted
}

func TestDeleteCity(t *testing.T) {
	var (
		req = weather.Request{
			City: "London",
		}

		cases = []database.TestCase{
			{
				Driver:  deleteCitySuccessPgxDriver{},
				Request: req,
				Iface:   test.New(nil, deleteCitySuccess, nil),
			},

			{
				Driver: deleteCityErrPgxDriver{},
				Request: weather.Request{
					City: "wdwadwadw",
				},
				Iface: test.New(weather.ErrNotDeleted, deleteCityErr, nil),
			},

			{
				Driver: deleteCityDbErrPgxDriver{},
				Request: weather.Request{
					City: "wdwadwadw",
				},
				Iface: test.New(weather.ErrNotDeleted, deleteCityDbErr, nil),
			},
		}
	)
	for _, c := range cases {
		ctx.Weather().SetDriver(c.Driver)
		r := c.Request.(weather.Request)
		err := ctx.Weather().Delete(r)
		c.Iface.Compare(t, nil, err)
	}
}

package database

import (
	"errors"
	"go.uber.org/zap"
	"projectname/internal/project/infrastructure/database/driver"
	"projectname/internal/project/infrastructure/database/driver/pgx"
	"projectname/pkg/test"
	"time"
)

const (
	DriverError   = `Driver Error`
	InvalidDriver = `Invalid Driver`
)

type (
	PgxDriverWithError struct{ pgx.Driver }

	TestCase struct {
		Driver  driver.Driver
		Request interface{}
		Iface   test.CaseInterface
	}
)

var (
	log              = zap.L()
	errDriverFailure = errors.New(`tcs.database.driver_error`)
	currentTime, _   = time.Parse(`"2006-01-02 15:04:05.999999"`,
		time.Now().UTC().Format(`"2006-01-02 15:04:05.999999"`))
)

func (d PgxDriverWithError) Get(_ interface{}, _ string, _ ...interface{}) error {
	return errDriverFailure
}

func (d PgxDriverWithError) Select(_ interface{}, _ string, _ ...interface{}) error {
	return errDriverFailure
}

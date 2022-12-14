package weather

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"projectname/internal/project/core/weather"
	domain "projectname/internal/project/domain/weather"
	"projectname/internal/project/interfaces/web/echo/middleware"
)

func Get(context echo.Context) error {
	var (
		ctx = context.(*middleware.Context)
	)

	res, err := weather.Get(ctx.Container)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, res)
}

func GetAvgTemp(context echo.Context) error {
	var (
		ctx = context.(*middleware.Context)
		req = domain.Request{
			City: context.Param(`name`),
		}
	)

	res, err := weather.GetAvgTemperature(ctx.Container, req)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, res)
}

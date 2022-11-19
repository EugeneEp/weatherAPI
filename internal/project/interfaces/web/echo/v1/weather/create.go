package weather

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"projectname/internal/project/core/weather"
	domain "projectname/internal/project/domain/weather"
	"projectname/internal/project/interfaces/web/echo/middleware"
)

func Create(context echo.Context) error {
	var (
		ctx = context.(*middleware.Context)
		req = domain.Request{
			City: context.Param(`name`),
		}
	)

	err := weather.Create(ctx.Container, req)
	if err != nil {
		return err
	}

	return context.NoContent(http.StatusOK)
}

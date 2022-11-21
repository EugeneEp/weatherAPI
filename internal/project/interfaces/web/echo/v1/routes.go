package v1

import (
	"github.com/labstack/echo/v4"
	"projectname/internal/project/interfaces/web/echo/v1/weather"
)

func Bind(g *echo.Group) {
	route := g.Group("/v1")
	weather.Bind(route)
}

package echo

import (
	"github.com/labstack/echo/v4"
	v4 "projectname/internal/project/interfaces/web/echo/v1"
)

func Bind(g *echo.Group) {
	route := g.Group("/api")
	v4.Bind(route)
}

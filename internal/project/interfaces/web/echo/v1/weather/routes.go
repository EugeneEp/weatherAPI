package weather

import (
	"github.com/labstack/echo/v4"
)

func Bind(g *echo.Group) {
	route := g.Group("/weather")
	route.GET("/city", Get)
	route.GET("/city/:name", Create)
	route.GET("/city/delete/:name", Delete)
}

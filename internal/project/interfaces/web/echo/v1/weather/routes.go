package weather

import (
	"github.com/labstack/echo/v4"
)

func Bind(g *echo.Group) {
	route := g.Group("/weather")
	route.GET("/city", Get)
	route.POST("/city/:name", Create)
	route.DELETE("/city/:name", Delete)
	route.GET("/city/:name", GetAvgTemp)
}

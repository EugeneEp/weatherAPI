package transport

import (
	"github.com/labstack/echo/v4"
)

// BrowserCachingOff Функция отключает кеширование
func BrowserCachingOff() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
			ctx.Response().Header().Set("cache-control", "no-store")
			ctx.Response().Header().Set("pragma", "no-cache")

			return next(ctx)
		}
	}
}

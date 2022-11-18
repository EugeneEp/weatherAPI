package server

import (
	"github.com/brpaz/echozap"
	framework "github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/sarulabs/di/v2"
	"go.uber.org/zap"
	"projectname/internal/project/interfaces/web/echo"
	"projectname/internal/project/interfaces/web/echo/middleware"
	"projectname/internal/project/interfaces/web/echo/middleware/transport"
)

const ServiceName = `EchoWebService`

func New(ctn di.Container, log *zap.Logger) (*framework.Echo, error) {
	web := framework.New()

	web.HideBanner = true
	web.HidePort = true
	web.Debug = true

	web.Use(transport.BrowserCachingOff())
	web.Use(echoMiddleware.RequestIDWithConfig(echoMiddleware.DefaultRequestIDConfig))
	web.Use(echoMiddleware.RecoverWithConfig(echoMiddleware.DefaultRecoverConfig))
	web.Use(echozap.ZapLogger(log))
	web.Use(middleware.WebContext(ctn))
	web.Pre(echoMiddleware.RemoveTrailingSlash())

	route := web.Group("")

	echo.Bind(route)

	return web, nil
}

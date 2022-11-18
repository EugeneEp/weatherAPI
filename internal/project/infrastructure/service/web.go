package service

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net"
	"net/http"
	"projectname/internal/project/domain/configuration"
	"projectname/internal/project/infrastructure/config"
	"projectname/internal/project/infrastructure/logger"
	"projectname/internal/project/interfaces/web/echo/server"
	"strconv"
)

func (a *app) startWeb() error {
	var log *zap.Logger

	if err := a.ctn.Fill(logger.BaseServiceName, &log); err != nil {
		return err
	}

	var cfg *viper.Viper

	if err := a.ctn.Fill(config.ServiceName, &cfg); err != nil {
		return err
	}

	var srv *echo.Echo

	if err := a.ctn.Fill(server.ServiceName, &srv); err != nil {
		return err
	}

	var listen string
	var tryListen string

	host := cfg.GetString(configuration.HttpHost)

	log.Info(`Base Web Server Configs`,
		zap.String(`host`, host),
		zap.String(`port`, cfg.GetString(configuration.HttpPort)),
	)

	for {
		port := cfg.GetString(configuration.HttpPort)

		if host == `` {
			tryListen = `localhost` + `:` + port
			listen = `:` + port
		} else {
			tryListen = host + `:` + port
			listen = tryListen
		}

		log.Info(`Trying to start Web Server`, zap.String(`addr`, listen))

		conn, err := net.Listen("tcp", tryListen)

		if err != nil {
			intPort, _ := strconv.Atoi(port)

			if intPort >= 65535 {
				intPort = 0
			} else {
				intPort++
			}

			cfg.Set(configuration.HttpPort, strconv.Itoa(intPort))
			continue
		}

		if conn != nil {
			_ = conn.Close()
			break
		}
	}

	log.Info(`Starting Web Server on ` + listen)

	if err := srv.Start(listen); err != nil {
		if err == http.ErrServerClosed {
			a.logger.Info("app.startWeb: Server stopped!")
			return err
		}
		return err
	}
	return nil
}

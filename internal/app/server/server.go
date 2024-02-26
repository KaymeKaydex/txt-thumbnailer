package server

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/KaymeKaydex/txt-thumbnailer/internal/app/config"
	"github.com/KaymeKaydex/txt-thumbnailer/internal/app/server/endpoints/convert"
)

func StartServer(ctx context.Context) error {
	cfg := config.FromContext(ctx)
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:     true,
		LogStatus:  true,
		LogLatency: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(log.Fields{
				"uri":     values.URI,
				"status":  values.Status,
				"latency": values.Latency,
			}).Info("request done")

			return nil
		},
	}))

	{
		convertController := convert.NewController(ctx)
		{
			e.POST("/convert", convertController.ToTxt)
		}
	}

	log.Info("available routes:")
	for _, route := range e.Routes() {
		log.Infof("%s %s", route.Method, route.Path)
	}

	log.Infof("starting http server at: %s", cfg.ListenerConfig.Address)
	return e.Start(cfg.ListenerConfig.Address)
}

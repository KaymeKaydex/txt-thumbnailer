package server

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/KaymeKaydex/txt-thumbnailer/internal/app/server/endpoints/convert"
)

func StartServer() error {
	e := echo.New()

	convertController := convert.NewController()

	e.POST("/convert", convertController.ToTxt)

	for _, route := range e.Routes() {
		log.Println(route.Method, route.Path)
	}

	return e.Start(":1323")
}

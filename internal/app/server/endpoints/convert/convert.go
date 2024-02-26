package convert

import (
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/KaymeKaydex/txt-thumbnailer/internal/app/converter"
	"github.com/KaymeKaydex/txt-thumbnailer/internal/app/server/models"
)

func (c *Controller) ToTxt(eCtx echo.Context) error {
	var err error

	if eCtx.Request().Body == nil {
		return eCtx.JSON(http.StatusBadRequest, models.BadResponse{Error: "no file in request"})
	}

	if eCtx.Request().Header.Get("Content-Type") != "text/plain" {
		return eCtx.JSON(http.StatusBadRequest, models.BadResponse{Error: "bad content-type"})
	}

	// get fontSize
	fontSizeInt := 20
	fontSize := eCtx.QueryParam("font-size")
	if fontSize != "" {
		fontSizeInt, err = strconv.Atoi(fontSize)
		if err != nil {
			log.Errorf("invalid font-size query param: %s", err)

			return err
		}
	}

	// get height
	heightInt := 1100
	height := eCtx.QueryParam("height")
	if height != "" {
		heightInt, err = strconv.Atoi(height)
		if err != nil {
			log.Errorf("invalid hight query param: %s", err)

			return err
		}
	}

	// get width
	widthInt := 700
	width := eCtx.QueryParam("width")
	if width != "" {
		widthInt, err = strconv.Atoi(width)
		if err != nil {
			log.Errorf("invalid width query param: %s", err)

			return err
		}
	}

	image, err := converter.Convert(converter.ConvertConfig{
		File:          eCtx.Request().Body,
		Font:          goregular.TTF,
		Height:        heightInt,
		Width:         widthInt,
		FontSize:      fontSizeInt,
		LineSpacing:   2,
		AutoEscape:    true,
		PaddingLeft:   0,
		PaddingTop:    0,
		PaddingBottom: 0,
		PaddingRight:  0,
	})
	if err != nil {
		log.Errorf("cant convert image with err: %s", err)

		return err
	}

	imgBts, err := io.ReadAll(image)
	if err != nil {
		return err
	}

	return eCtx.Blob(http.StatusOK, "image/jpeg", imgBts)
}

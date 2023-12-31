package converter

import (
	"bytes"
	"image"
	"image/draw"
	"image/jpeg"
	"io"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type ConvertConfig struct {
	File          io.Reader
	Height        int
	Width         int
	FontPath      string
	FontSize      int
	AutoEscape    bool
	PaddingLeft   int
	LineSpacing   int
	PaddingTop    int
	PaddingRight  int
	PaddingBottom int
	Font          []byte
}

func Convert(cfg ConvertConfig) (*bytes.Buffer, error) {
	// Читаем содержимое текстового файла
	content, err := io.ReadAll(cfg.File)
	if err != nil {
		log.Errorf("cant read file from cfg.File with err: %s", err)

		return nil, err
	}

	// Создаем новое изображение для миниатюры
	thumbnail := image.NewRGBA(image.Rect(-cfg.PaddingLeft, -cfg.PaddingTop, cfg.Width, cfg.Height))
	draw.Draw(thumbnail, thumbnail.Bounds(), &image.Uniform{C: image.White}, image.Point{}, draw.Src)

	// Инициализируем контекст рендеринга FreeType
	fnt, err := freetype.ParseFont(cfg.Font)
	if err != nil {
		return nil, err
	}

	fntContext := freetype.NewContext()
	fntContext.SetDPI(72)
	fntContext.SetFont(fnt)
	fntContext.SetFontSize(float64(cfg.FontSize))
	fntContext.SetClip(thumbnail.Bounds())
	fntContext.SetDst(thumbnail)
	fntContext.SetSrc(image.Black)
	face := truetype.NewFace(fnt, &truetype.Options{
		Size: float64(cfg.FontSize),
		DPI:  72,
	})

	// Разбиваем текст на строки и рендерим каждую строку
	lines := bytes.Split(content, []byte("\n"))
	y := fixed.I(cfg.FontSize).Ceil()
	for _, line := range lines {
		if y > cfg.Height-cfg.PaddingBottom {
			break
		}

		// Проверяем, помещается ли строка в оставшееся пространство изображения
		metrics := font.MeasureString(face, string(line))
		if metrics.Ceil() > cfg.Width-cfg.PaddingRight-10 { // todo -10 это костыль
			// Если строка не помещается, разбиваем ее на слова и пытаемся перенести слова на следующую строку
			words := bytes.Split(line, []byte(" "))
			var currentLine []byte
			for i, word := range words {
				testLine := append(currentLine, word...)
				metrics := font.MeasureString(face, string(testLine))
				if metrics.Ceil() > cfg.Width-cfg.PaddingRight-10 { // todo -10 это костыль
					// Строка не помещается, рендерим предыдущую и начинаем новую строку
					fntContext.DrawString(string(currentLine), fixed.Point26_6{X: 0, Y: fixed.I(y)})
					y += fixed.I(cfg.FontSize + cfg.LineSpacing).Ceil()
					currentLine = word
				} else {
					// Слово помещается на текущей строке
					if i != 0 {
						currentLine = append(currentLine, ' ')
					}

					currentLine = append(currentLine, word...)
				}
			}
			// Рендерим оставшуюся строку
			fntContext.DrawString(string(currentLine), fixed.Point26_6{X: 0, Y: fixed.I(y)})
			y += fixed.I(cfg.FontSize + cfg.LineSpacing).Ceil()
		} else {
			// Строка помещается в оставшееся пространство на текущей строке
			fntContext.DrawString(string(line), fixed.Point26_6{X: 0, Y: fixed.I(y)})
			y += fixed.I(cfg.FontSize + cfg.LineSpacing).Ceil()
		}
	}

	b := bytes.NewBuffer([]byte{})

	jpeg.Encode(b, thumbnail, nil)

	return b, nil
}

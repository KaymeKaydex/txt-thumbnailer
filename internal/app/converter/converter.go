package converter

import (
	"image"
	"image/draw"
	"log"
	"os"
)

type ConvertConfig struct {
	Height      int
	Width       int
	Out         string
	FontPath    string
	FontSize    uint64
	AutoEscape  bool
	Padding     int
	LineSpacing uint64
}

func Convert(cfg ConvertConfig) {
	// Читаем содержимое текстового файла
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// Создаем новое изображение для миниатюры
	thumbnail := image.NewRGBA(image.Rect(cfg.Padding, cfg.Padding, cfg.Width, cfg.Height))
	draw.Draw(thumbnail, thumbnail.Bounds(), &image.Uniform{C: image.White}, image.Point{}, draw.Src)
}

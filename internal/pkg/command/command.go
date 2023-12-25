package command

import (
	"bytes"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/KaymeKaydex/txt-thumbnailer/internal/app/converter"
	"github.com/KaymeKaydex/txt-thumbnailer/internal/app/server"
)

var Commands = []*cobra.Command{
	cmdConvert(),
	cmdServer(),
}

func cmdConvert() *cobra.Command {
	convertCommand := &cobra.Command{
		Use:     "convert [path to txt file to convert]",
		Short:   "convert any txt file to jpg image",
		Long:    `Convert command can convert any txt file to jpg.`,
		Args:    cobra.MinimumNArgs(1),
		Example: "$ go run cmd/txt-thumbnailer/main.go convert examples/txt/long.txt  --font-size=16  --padding-left=50 --padding-top=50 --padding-right=50 --padding-bottom=50 --font=examples/fonts/MailSansRoman-Light.ttf",
	}

	height := convertCommand.PersistentFlags().Int("height", 1100, "height of result image")
	width := convertCommand.PersistentFlags().Int("width", 700, "width of result image")
	out := convertCommand.PersistentFlags().String("out", "result.jpg", "~/myimage.jpg")
	fontPath := convertCommand.PersistentFlags().String("font", "", "~/font.ttf , default font is goregular.TTF")
	fontSize := convertCommand.PersistentFlags().Int("font-size", 20, "font size for txt symbols")
	convertCommand.PersistentFlags().Bool("auto-escape", true, "escapes your txt thumbnail file lines if u need")
	lineSpacing := convertCommand.PersistentFlags().Int("line-spacing", 2, "space between lines")

	paddingLeft := convertCommand.PersistentFlags().Int("padding-left", 0, "padding form left border in pixels")
	paddingTop := convertCommand.PersistentFlags().Int("padding-top", 0, "padding form top border in pixels")
	paddingRight := convertCommand.PersistentFlags().Int("padding-right", 0, "padding form top border in pixels")
	paddingBottom := convertCommand.PersistentFlags().Int("padding-bottom", 0, "padding form top border in pixels")

	convertCommand.Run = func(cmd *cobra.Command, args []string) {
		t := time.Now()
		file, err := os.ReadFile(args[0])
		if err != nil {
			log.Fatalln(err)
		}

		cfg := converter.ConvertConfig{
			Height:        *height,
			Width:         *width,
			FontSize:      *fontSize,
			LineSpacing:   *lineSpacing,
			File:          bytes.NewBuffer(file),
			PaddingLeft:   *paddingLeft,
			PaddingTop:    *paddingTop,
			PaddingRight:  *paddingRight,
			PaddingBottom: *paddingBottom,
			Font:          goregular.TTF,
		}

		if *fontPath != "" {
			fontBytes, err := os.ReadFile(*fontPath)
			if err != nil {
				log.Fatal(err)
			}
			cfg.Font = fontBytes
		}

		res, err := converter.Convert(cfg)
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("Thumbnail generated successfully! %s\n", time.Since(t))

		// Сохраняем миниатюру в файловую систему
		err = os.WriteFile(*out, res.Bytes(), 0644)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return convertCommand
}

func cmdServer() *cobra.Command {
	serverCommand := &cobra.Command{
		Use:   "server [command for server]",
		Short: "Start txt-thumbnailer server",
		Long:  `Start txt-thumbnailer server. By default config file configs/config.yaml`,
		Args:  cobra.MinimumNArgs(0),
	}

	cfg := serverCommand.PersistentFlags().String("config", "configs/config.yaml", "config file for server")
	if cfg == nil || *cfg == "" {
		log.Fatalln("cant get config param from args")
	}

	serverCommand.Run = func(cmd *cobra.Command, args []string) {
		err := server.StartServer()
		if err != nil {
			log.Fatalln(err)
		}
	}

	return serverCommand
}

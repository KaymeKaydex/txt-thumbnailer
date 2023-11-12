package command

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var Commands = []*cobra.Command{
	cmdConvert(),
	cmdServer(),
}

func cmdConvert() *cobra.Command {
	convertCommand := &cobra.Command{
		Use:   "convert [path to txt file to convert]",
		Short: "convert any txt file to jpg image",
		Long:  `Convert command can convert any txt file to jpg.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}

	convertCommand.PersistentFlags().Uint("height", 3508, "height of result image")
	convertCommand.PersistentFlags().Uint("width", 2480, "width of result image")
	convertCommand.PersistentFlags().String("out", "result.jpg", "~/myimage.jpg")
	convertCommand.PersistentFlags().String("font", "font.ttf", "~/font.ttf")
	convertCommand.PersistentFlags().Uint("font-size", 20, "font size for txt symbols")
	convertCommand.PersistentFlags().Bool("auto-escape", true, "escapes your txt thumbnail file lines if u need")
	convertCommand.PersistentFlags().Uint("padding", 10, "padding form border in pixels")
	convertCommand.PersistentFlags().Uint("line-spacing", 2, "space between lines")

	return convertCommand
}

func cmdServer() *cobra.Command {
	return &cobra.Command{
		Use:   "server [string to print]",
		Short: "[NOW IS NOT AVAILABLE] Print anything to the screen",
		Long: `[NOW IS NOT AVAILABLE] print is for printing anything back to the screen.
For many years people have printed back to the screen.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}
}

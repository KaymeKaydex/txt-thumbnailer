package main

import (
	"github.com/spf13/cobra"

	"github.com/KaymeKaydex/txt-thumbnailer/internal/pkg/command"
)

var rootCmd = &cobra.Command{Use: "txt-thumbnailer"}

func main() {
	rootCmd.AddCommand(command.Commands...)
	rootCmd.Execute()
}

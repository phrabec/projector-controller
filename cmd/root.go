package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "Panasonic PT-EX610/EX610L projector controller",
}

var (
	device      string
	projectorId string
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&device, "device", "/dev/ttyUSB0", "Serial port device")
	rootCmd.PersistentFlags().StringVar(&projectorId, "projector-id", kIdAll, "Target projector ID")

	rootCmd.AddCommand(powerCmd)
	rootCmd.AddCommand(shutterCmd)
	rootCmd.AddCommand(inputCmd)
	rootCmd.AddCommand(aspectCmd)
	rootCmd.AddCommand(customCmd)
}

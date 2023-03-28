package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "Panasonic PT-EX610/EX610L projector controller",
}

var (
	device      string
	projectorID string
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&device, "device", "/dev/ttyUSB0", "Serial port device")
	rootCmd.PersistentFlags().StringVar(&projectorID, "projector-id", idAll, "Target projector ID")

	rootCmd.AddCommand(powerCmd)
	rootCmd.AddCommand(shutterCmd)
	rootCmd.AddCommand(inputCmd)
	rootCmd.AddCommand(aspectCmd)
	rootCmd.AddCommand(customCmd)
}

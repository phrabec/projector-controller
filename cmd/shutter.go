package cmd

import (
	"github.com/spf13/cobra"
)

var shutterCmd = &cobra.Command{
	Use:   "shutter",
	Short: "Get or set shutter state",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorID, commandShutterStatus, paramsNone)
	},
}

var shutterOnCmd = &cobra.Command{
	Use:   "open",
	Short: "Open shutter",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorID, commandShutter, paramsShutterOn)
	},
}

var shutterOffCmd = &cobra.Command{
	Use:   "close",
	Short: "Close shutter",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorID, commandShutter, paramsShutterOff)
	},
}

func init() {
	shutterCmd.AddCommand(shutterOnCmd)
	shutterCmd.AddCommand(shutterOffCmd)
}

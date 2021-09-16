package cmd

import (
	"github.com/spf13/cobra"
)

var shutterCmd = &cobra.Command{
	Use:   "shutter",
	Short: "Get or set shutter state",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorId, kCommandShutterStatus, kParamsNone)
	},
}

var shutterOnCmd = &cobra.Command{
	Use:   "open",
	Short: "Open shutter",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorId, kCommandShutter, kParamsShutterOn)
	},
}

var shutterOffCmd = &cobra.Command{
	Use:   "close",
	Short: "Close shutter",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorId, kCommandShutter, kParamsShutterOff)
	},
}

func init() {
	shutterCmd.AddCommand(shutterOnCmd)
	shutterCmd.AddCommand(shutterOffCmd)
}

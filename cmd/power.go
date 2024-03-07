package cmd

import (
	"github.com/spf13/cobra"
)

var powerCmd = &cobra.Command{
	Use:   "power",
	Short: "Get or set power state",
	RunE: func(_ *cobra.Command, _ []string) error {
		return communicate(device, projectorID, commandMainPowerStatus, paramsNone)
	},
}

var powerOnCmd = &cobra.Command{
	Use:   "on",
	Short: "Power on projector",
	RunE: func(_ *cobra.Command, _ []string) error {
		return communicate(device, projectorID, commandPowerOn, paramsNone)
	},
}

var powerOffCmd = &cobra.Command{
	Use:   "off",
	Short: "Power off projector",
	RunE: func(_ *cobra.Command, _ []string) error {
		return communicate(device, projectorID, commandPowerOff, paramsNone)
	},
}

func init() {
	powerCmd.AddCommand(powerOnCmd)
	powerCmd.AddCommand(powerOffCmd)
}

package cmd

import (
	"github.com/spf13/cobra"
)

var powerCmd = &cobra.Command{
	Use:   "power",
	Short: "Get or set power state",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorId, kCommandMainPowerStatus, kParamsNone)
	},
}

var powerOnCmd = &cobra.Command{
	Use:   "on",
	Short: "Power on projector",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorId, kCommandPowerOn, kParamsNone)
	},
}

var powerOffCmd = &cobra.Command{
	Use:   "off",
	Short: "Power off projector",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorId, kCommandPowerOff, kParamsNone)
	},
}

func init() {
	powerCmd.AddCommand(powerOnCmd)
	powerCmd.AddCommand(powerOffCmd)
}

package cmd

import (
	"github.com/spf13/cobra"
)

var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "Set input configutation",
}

var inputDigitalLinkCmd = &cobra.Command{
	Use:   "digital-link",
	Short: "Set input to DigitalLink",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorId, kCommandInputSelection, kParamsDigitalLink)
	},
}

var inputDisplaPort = &cobra.Command{
	Use:   "display-port",
	Short: "Set input to DisplayPort",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorId, kCommandInputSelection, kParamsDisplayPort)
	},
}

var inputHdmiCmd = &cobra.Command{
	Use:   "hdmi",
	Short: "Set input to HDMI",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorId, kCommandInputSelection, kParamsHdmi)
	},
}

var inputDviCmd = &cobra.Command{
	Use:   "dvi",
	Short: "Set input to DVI-D",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorId, kCommandInputSelection, kParamsDvi)
	},
}

var inputRgb1Cmd = &cobra.Command{
	Use:   "rgb1",
	Short: "Set input to RGB 1",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorId, kCommandInputSelection, kParamsRgb1)
	},
}

var inputRgb2Cmd = &cobra.Command{
	Use:   "rgb2",
	Short: "Set input to RGB 2",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorId, kCommandInputSelection, kParamsRgb2)
	},
}

var inputVideoCmd = &cobra.Command{
	Use:   "video",
	Short: "Set input to Video",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorId, kCommandInputSelection, kParamsVideo)
	},
}

func init() {
	inputCmd.AddCommand(inputDigitalLinkCmd)
	inputCmd.AddCommand(inputDisplaPort)
	inputCmd.AddCommand(inputHdmiCmd)
	inputCmd.AddCommand(inputDviCmd)
	inputCmd.AddCommand(inputRgb1Cmd)
	inputCmd.AddCommand(inputRgb2Cmd)
	inputCmd.AddCommand(inputVideoCmd)
}

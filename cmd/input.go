package cmd

import (
	"github.com/spf13/cobra"
)

var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "Get or set input configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorID, commandInputSignalStatus, paramsNone)
	},
}

var inputDigitalLinkCmd = &cobra.Command{
	Use:   "digital-link",
	Short: "Set input to DigitalLink",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorID, commandInputSelection, paramsDigitalLink)
	},
}

var inputDisplaPort = &cobra.Command{
	Use:   "display-port",
	Short: "Set input to DisplayPort",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorID, commandInputSelection, paramsDisplayPort)
	},
}

var inputHdmiCmd = &cobra.Command{
	Use:   "hdmi",
	Short: "Set input to HDMI",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorID, commandInputSelection, paramsHdmi)
	},
}

var inputDviCmd = &cobra.Command{
	Use:   "dvi",
	Short: "Set input to DVI-D",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorID, commandInputSelection, paramsDvi)
	},
}

var inputRgb1Cmd = &cobra.Command{
	Use:   "rgb1",
	Short: "Set input to RGB 1",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorID, commandInputSelection, paramsRgb1)
	},
}

var inputRgb2Cmd = &cobra.Command{
	Use:   "rgb2",
	Short: "Set input to RGB 2",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorID, commandInputSelection, paramsRgb2)
	},
}

var inputVideoCmd = &cobra.Command{
	Use:   "video",
	Short: "Set input to Video",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorID, commandInputSelection, paramsVideo)
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

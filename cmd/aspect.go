package cmd

import (
	"github.com/spf13/cobra"
)

var aspectCmd = &cobra.Command{
	Use:   "aspect",
	Short: "Get or set aspect ratio",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorId, kCommandAspectModeStatus, kParamsNone)
	},
}

var aspectAutoCmd = &cobra.Command{
	Use:   "auto",
	Short: "Set aspect ratio to auto",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorId, kCommandAspect, kParamsAspectAuto)
	},
}

var aspect43Cmd = &cobra.Command{
	Use:   "43",
	Short: "Set aspect ratio to 4:3",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorId, kCommandAspect, kParamsAspectAuto)
	},
}

var aspectWideCmd = &cobra.Command{
	Use:   "wide",
	Short: "Set aspect ratio to wide",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorId, kCommandAspect, kParamsAspectAuto)
	},
}

var aspectNativeCmd = &cobra.Command{
	Use:   "native",
	Short: "Set aspect ratio to native",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorId, kCommandAspect, kParamsAspectAuto)
	},
}

var aspectFullCmd = &cobra.Command{
	Use:   "full",
	Short: "Set aspect ratio to full",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorId, kCommandAspect, kParamsAspectAuto)
	},
}

var aspectHorizontalFitCmd = &cobra.Command{
	Use:   "h-fit",
	Short: "Set aspect ratio to horizontal fit",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorId, kCommandAspect, kParamsAspectAuto)
	},
}

var aspectVerticalFitCmd = &cobra.Command{
	Use:   "v-fit",
	Short: "Set aspect ratio to vertical fit",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorId, kCommandAspect, kParamsAspectAuto)
	},
}

func init() {
	aspectCmd.AddCommand(aspectAutoCmd)
	aspectCmd.AddCommand(aspect43Cmd)
	aspectCmd.AddCommand(aspectWideCmd)
	aspectCmd.AddCommand(aspectNativeCmd)
	aspectCmd.AddCommand(aspectFullCmd)
	aspectCmd.AddCommand(aspectHorizontalFitCmd)
	aspectCmd.AddCommand(aspectVerticalFitCmd)
}

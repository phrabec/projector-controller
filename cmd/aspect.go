package cmd

import (
	"github.com/spf13/cobra"
)

var aspectCmd = &cobra.Command{
	Use:   "aspect",
	Short: "Get or set aspect ratio",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorID, commandAspectModeStatus, paramsNone)
	},
}

var aspectAutoCmd = &cobra.Command{
	Use:   "auto",
	Short: "Set aspect ratio to auto",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorID, commandAspect, paramsAspectAuto)
	},
}

var aspect43Cmd = &cobra.Command{
	Use:   "43",
	Short: "Set aspect ratio to 4:3",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorID, commandAspect, paramsAspect43)
	},
}

var aspectWideCmd = &cobra.Command{
	Use:   "wide",
	Short: "Set aspect ratio to wide",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorID, commandAspect, paramsAspectWide)
	},
}

var aspectRealCmd = &cobra.Command{
	Use:   "real",
	Short: "Set aspect ratio to real",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorID, commandAspect, paramsAspectReal)
	},
}

var aspectFullCmd = &cobra.Command{
	Use:   "full",
	Short: "Set aspect ratio to full",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorID, commandAspect, paramsAspectFull)
	},
}

var aspectHorizontalFitCmd = &cobra.Command{
	Use:   "h-fit",
	Short: "Set aspect ratio to horizontal fit",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorID, commandAspect, paramsAspectHFit)
	},
}

var aspectVerticalFitCmd = &cobra.Command{
	Use:   "v-fit",
	Short: "Set aspect ratio to vertical fit",
	RunE: func(cmd *cobra.Command, args []string) error {
		return communicate(device, projectorID, commandAspect, paramsAspectVFit)
	},
}

func init() {
	aspectCmd.AddCommand(aspectAutoCmd)
	aspectCmd.AddCommand(aspect43Cmd)
	aspectCmd.AddCommand(aspectWideCmd)
	aspectCmd.AddCommand(aspectRealCmd)
	aspectCmd.AddCommand(aspectFullCmd)
	aspectCmd.AddCommand(aspectHorizontalFitCmd)
	aspectCmd.AddCommand(aspectVerticalFitCmd)
}

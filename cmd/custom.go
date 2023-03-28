package cmd

import (
	"github.com/spf13/cobra"
)

var customCmd = &cobra.Command{
	Use:   "custom [flags] command [params]",
	Short: "Send custom command to projector",
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := paramsNone
		if len(args) > 1 {
			params = args[1]
		}
		return communicate(device, projectorID, args[0], params)
	},
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.bug.st/serial"
)

var serialCmd = &cobra.Command{
	Use:   "serial",
	Short: "Serial port tools",
}

var listPortsCmd = &cobra.Command{
	Use:   "list",
	Short: "List available serial ports",
	RunE: func(_ *cobra.Command, _ []string) error {
		ports, err := serial.GetPortsList()
		if err != nil {
			return err
		}

		fmt.Printf("Found %d serial ports:\n", len(ports))
		for _, port := range ports {
			fmt.Println(port)
		}

		return nil
	},
}

func init() {
	serialCmd.AddCommand(listPortsCmd)
}

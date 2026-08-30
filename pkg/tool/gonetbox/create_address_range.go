package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createAddressRange(c *client.Client) *cobra.Command {
	var status string
	var description string
	result := &cobra.Command{
		Use:   "create-address-range [start] [end]",
		Short: "Create an IP address range (addresses in CIDR notation)",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(
				c.CreateAddressRange(
					arguments[0],
					arguments[1],
					status,
					description,
				),
			)
		},
	}
	result.Flags().StringVar(
		&status,
		"status",
		"",
		"range status (active, reserved, deprecated)",
	)
	result.Flags().StringVar(
		&description,
		"description",
		"",
		"range description",
	)

	return result
}

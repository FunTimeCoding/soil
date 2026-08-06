package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createAddress(c *client.Client) *cobra.Command {
	var interfaceName string
	var status string
	result := &cobra.Command{
		Use:   "create-address [device] [address]",
		Short: "Assign an IP address to a device interface",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(
				c.CreateAddress(
					arguments[0],
					interfaceName,
					arguments[1],
					status,
				),
			)
		},
	}
	result.Flags().StringVar(
		&interfaceName,
		"interface",
		"",
		"interface name (required)",
	)
	result.Flags().StringVar(
		&status,
		"status",
		"",
		"address status (active, dhcp, reserved)",
	)
	errors.PanicOnError(result.MarkFlagRequired("interface"))

	return result
}

package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createDeviceTunnelTermination(c *client.Client) *cobra.Command {
	var tunnel string
	var interfaceName string
	var role string
	result := &cobra.Command{
		Use:   "create-device-tunnel-termination [device]",
		Short: "Create a tunnel termination on a device interface",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(
				c.CreateDeviceTunnelTermination(
					arguments[0],
					tunnel,
					interfaceName,
					role,
				),
			)
		},
	}
	result.Flags().StringVar(
		&tunnel,
		"tunnel",
		"",
		"tunnel name (required)",
	)
	result.Flags().StringVar(
		&interfaceName,
		"interface",
		"",
		"interface name on the device (required)",
	)
	result.Flags().StringVar(
		&role,
		"role",
		"",
		"termination role: peer, hub, spoke (required)",
	)
	errors.PanicOnError(result.MarkFlagRequired("tunnel"))
	errors.PanicOnError(result.MarkFlagRequired("interface"))
	errors.PanicOnError(result.MarkFlagRequired("role"))

	return result
}

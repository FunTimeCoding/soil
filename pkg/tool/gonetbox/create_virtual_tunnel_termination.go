package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createVirtualTunnelTermination(c *client.Client) *cobra.Command {
	var tunnel string
	var interfaceName string
	var role string
	result := &cobra.Command{
		Use:   "create-virtual-tunnel-termination [vm]",
		Short: "Create a tunnel termination on a VM interface",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(
				c.CreateVirtualTunnelTermination(
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
		"interface name on the VM (required)",
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

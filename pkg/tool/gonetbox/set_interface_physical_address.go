package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func setInterfacePhysicalAddress(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "set-interface-physical-address [device] [interface] [address]",
		Short: "Assign a MAC address to an existing device interface",
		Args:  cobra.ExactArgs(3),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(
				c.SetInterfacePhysicalAddress(
					arguments[0],
					arguments[1],
					arguments[2],
				),
			)
		},
	}
}

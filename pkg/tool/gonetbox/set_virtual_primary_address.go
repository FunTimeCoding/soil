package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	generated "github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/spf13/cobra"
)

func setVirtualPrimaryAddress(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "set-virtual-primary-address [vm] [address]",
		Short: "Set the primary IP address of a virtual machine",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(
				c.UpdateVirtualMachine(
					arguments[0],
					generated.UpdateVirtualMachineRequest{
						PrimaryAddress: &arguments[1],
					},
				),
			)
		},
	}
}

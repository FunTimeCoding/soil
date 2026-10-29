package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/spf13/cobra"
)

func setVirtualPrimaryAddress(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "set-virtual-primary-address [vm] [address]",
		Short: "Set the primary IP address of a virtual machine",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(
				c.UpdateVirtualMachine(
					context.Background(),
					arguments[0],
					client.UpdateVirtualMachineRequest{
						PrimaryAddress: &arguments[1],
					},
				),
			)
		},
	}
}

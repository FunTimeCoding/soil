package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func setInterfacePhysicalAddress(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "set-interface-physical-address [device] [interface] [address]",
		Short: "Assign a MAC address to an existing device interface",
		Args:  cobra.ExactArgs(3),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.Emit(
				c.SetInterfacePhysicalAddress(
					arguments[0],
					arguments[1],
					arguments[2],
				),
			)
		},
	}
}

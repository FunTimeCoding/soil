package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func setVirtualInterfacePhysicalAddress(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   constant.VirtualPhysicalAddressUsage,
		Short: "Assign a MAC address to an existing virtual machine interface",
		Args:  cobra.ExactArgs(3),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.Emit(
				c.SetVirtualInterfacePhysicalAddress(
					arguments[0],
					arguments[1],
					arguments[2],
				),
			)
		},
	}
}

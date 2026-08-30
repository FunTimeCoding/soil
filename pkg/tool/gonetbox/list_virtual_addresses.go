package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listVirtualAddresses(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-virtual-addresses [vm]",
		Short: "List IP addresses for a virtual machine",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(c.ListVirtualAddresses(arguments[0]))
		},
	}
}

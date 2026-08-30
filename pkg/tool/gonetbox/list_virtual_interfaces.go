package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listVirtualInterfaces(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-virtual-interfaces [vm]",
		Short: "List interfaces on a virtual machine",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(c.ListVirtualInterfaces(arguments[0]))
		},
	}
}

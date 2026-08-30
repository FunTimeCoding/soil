package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listVirtualDisks(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-virtual-disks [vm]",
		Short: "List disks on a virtual machine",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(c.ListVirtualDisks(arguments[0]))
		},
	}
}

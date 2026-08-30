package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listVirtualMachines(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-virtual-machines",
		Short: "List all NetBox virtual machines",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.ListVirtualMachines())
		},
	}
}

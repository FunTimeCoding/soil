package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listVirtualMachines(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "list-virtual-machines",
		Short: "List all NetBox virtual machines",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.Emit(c.ListVirtualMachines())
		},
	}
}

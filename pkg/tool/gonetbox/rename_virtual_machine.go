package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	generated "github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/spf13/cobra"
)

func renameVirtualMachine(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "rename-virtual-machine [name] [new-name]",
		Short: "Rename a NetBox virtual machine",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(
				c.UpdateVirtualMachine(
					arguments[0],
					generated.UpdateVirtualMachineRequest{Name: &arguments[1]},
				),
			)
		},
	}
}

package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/spf13/cobra"
)

func renameVirtualMachine(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "rename-virtual-machine [name] [new-name]",
		Short: "Rename a NetBox virtual machine",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(
				c.UpdateVirtualMachine(
					context.Background(),
					arguments[0],
					client.UpdateVirtualMachineRequest{Name: &arguments[1]},
				),
			)
		},
	}
}

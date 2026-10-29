package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func createDeviceRole(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var virtualMachine bool
	result := &cobra.Command{
		Use:   "create-device-role [name]",
		Short: "Create a NetBox device role",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(
				c.CreateDeviceRole(
					context.Background(),
					arguments[0],
					virtualMachine,
				),
			)
		},
	}
	result.Flags().BoolVar(
		&virtualMachine,
		"virtual-machine",
		false,
		"allow virtual machines to take this role",
	)

	return result
}

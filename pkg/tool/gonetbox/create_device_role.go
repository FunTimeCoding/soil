package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createDeviceRole(
	c *client.Client,
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
			t.Emit(c.CreateDeviceRole(arguments[0], virtualMachine))
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

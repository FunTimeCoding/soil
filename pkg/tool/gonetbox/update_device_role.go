package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/spf13/cobra"
)

func updateDeviceRole(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var name string
	result := &cobra.Command{
		Use:   "update-device-role [role]",
		Short: "Update a device role; a new name brings its slug along",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			body := client.UpdateDeviceRoleRequest{}

			if name != "" {
				body.Name = &name
			}

			t.EmitValue(
				c.UpdateDeviceRole(context.Background(), arguments[0], body),
			)
		},
	}
	result.Flags().StringVar(&name, "name", "", "new device role name")

	return result
}

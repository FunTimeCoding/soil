package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	generated "github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/spf13/cobra"
)

func renameDevice(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "rename-device [name] [new-name]",
		Short: "Rename a NetBox device",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.Emit(
				c.UpdateDevice(
					arguments[0],
					generated.UpdateDeviceRequest{Name: &arguments[1]},
				),
			)
		},
	}
}

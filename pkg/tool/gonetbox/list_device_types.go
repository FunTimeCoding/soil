package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listDeviceTypes(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "list-device-types",
		Short: "List all NetBox device types",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.Emit(c.ListDeviceTypes())
		},
	}
}

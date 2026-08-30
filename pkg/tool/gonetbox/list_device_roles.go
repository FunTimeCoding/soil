package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listDeviceRoles(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-device-roles",
		Short: "List all NetBox device roles",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.ListDeviceRoles())
		},
	}
}

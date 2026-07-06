package gonetbox

import (
	"fmt"
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
			fmt.Println(c.ListDeviceRoles())
		},
	}
}

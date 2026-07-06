package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listDeviceTypes(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-device-types",
		Short: "List all NetBox device types",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(c.ListDeviceTypes())
		},
	}
}

package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createDeviceRole(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "create-device-role [name]",
		Short: "Create a NetBox device role",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.CreateDeviceRole(arguments[0]))
		},
	}
}

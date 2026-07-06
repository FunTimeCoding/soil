package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func addDeviceTag(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "add-device-tag [device] [tag]",
		Short: "Add a tag to a device",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.AddDeviceTag(arguments[0], arguments[1]))
		},
	}
}

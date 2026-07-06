package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func removeDeviceTag(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "remove-device-tag [device] [tag]",
		Short: "Remove a tag from a device",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.RemoveDeviceTag(arguments[0], arguments[1]))
		},
	}
}

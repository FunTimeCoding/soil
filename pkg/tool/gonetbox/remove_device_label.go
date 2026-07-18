package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func removeDeviceLabel(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "remove-device-label [device] [key]",
		Short: "Remove a label from a device",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			c.RemoveDeviceLabel(arguments[0], arguments[1])
			fmt.Println("label removed")
		},
	}
}

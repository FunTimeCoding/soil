package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listDeviceLabels(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-device-labels [device]",
		Short: "List labels on a device",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.ListDeviceLabels(arguments[0]))
		},
	}
}

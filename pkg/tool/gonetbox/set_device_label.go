package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func setDeviceLabel(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "set-device-label [device] [key] [value]",
		Short: "Set a key-value label on a device",
		Args:  cobra.ExactArgs(3),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(
				c.SetDeviceLabel(
					arguments[0],
					arguments[1],
					arguments[2],
				),
			)
		},
	}
}

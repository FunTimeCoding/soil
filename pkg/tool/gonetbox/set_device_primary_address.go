package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	generated "github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/spf13/cobra"
)

func setDevicePrimaryAddress(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "set-device-primary-address [device] [address]",
		Short: "Set the primary IP address of a device",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(
				c.UpdateDevice(
					arguments[0],
					generated.UpdateDeviceRequest{
						PrimaryAddress: &arguments[1],
					},
				),
			)
		},
	}
}

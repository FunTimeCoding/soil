package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/spf13/cobra"
)

func setDevicePrimaryAddress(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "set-device-primary-address [device] [address]",
		Short: "Set the primary IP address of a device",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(
				c.UpdateDevice(
					context.Background(),
					arguments[0],
					client.UpdateDeviceRequest{PrimaryAddress: &arguments[1]},
				),
			)
		},
	}
}

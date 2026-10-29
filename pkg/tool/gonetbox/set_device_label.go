package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func setDeviceLabel(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "set-device-label [device] [key] [value]",
		Short: "Set a key-value label on a device",
		Args:  cobra.ExactArgs(3),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(
				c.SetDeviceLabel(
					context.Background(),
					arguments[0],
					arguments[1],
					arguments[2],
				),
			)
		},
	}
}

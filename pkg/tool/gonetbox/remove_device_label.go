package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func removeDeviceLabel(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "remove-device-label [device] [key]",
		Short: "Remove a label from a device",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitDone(
				c.RemoveDeviceLabel(
					context.Background(),
					arguments[0],
					arguments[1],
				),
			)
			console.Line("label removed")
		},
	}
}

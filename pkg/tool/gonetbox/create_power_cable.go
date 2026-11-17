package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func createPowerCable(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "create-power-cable [outlet-device] [outlet] [port-device] [port]",
		Short: "Create a power cable from a power outlet to a power port",
		Args:  cobra.ExactArgs(4),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(
				c.CreatePowerCable(
					context.Background(),
					arguments[0],
					arguments[1],
					arguments[2],
					arguments[3],
				),
			)
		},
	}
}

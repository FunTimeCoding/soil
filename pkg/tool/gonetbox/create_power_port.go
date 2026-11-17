package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func createPowerPort(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "create-power-port [device] [name]",
		Short: "Create a power port, the inlet a device is powered through",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(
				c.CreatePowerPort(
					context.Background(),
					arguments[0],
					arguments[1],
				),
			)
		},
	}
}

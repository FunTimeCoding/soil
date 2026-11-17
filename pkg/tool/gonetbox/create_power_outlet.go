package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func createPowerOutlet(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "create-power-outlet [device] [name]",
		Short: "Create a power outlet on a device, such as a battery or a power strip",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(
				c.CreatePowerOutlet(
					context.Background(),
					arguments[0],
					arguments[1],
				),
			)
		},
	}
}

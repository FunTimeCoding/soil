package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func createCable(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "create-cable [device-a] [interface-a] [device-b] [interface-b]",
		Short: "Create a NetBox cable between two device interfaces",
		Args:  cobra.ExactArgs(4),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(
				c.CreateCable(
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

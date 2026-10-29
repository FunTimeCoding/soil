package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func setVirtualLabel(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "set-virtual-label [machine] [key] [value]",
		Short: "Set a key-value label on a virtual machine",
		Args:  cobra.ExactArgs(3),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(
				c.SetVirtualLabel(
					context.Background(),
					arguments[0],
					arguments[1],
					arguments[2],
				),
			)
		},
	}
}

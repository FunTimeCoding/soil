package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func removeVirtualTag(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "remove-virtual-tag [vm] [tag]",
		Short: "Remove a tag from a virtual machine",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(
				c.RemoveVirtualTag(
					context.Background(),
					arguments[0],
					arguments[1],
				),
			)
		},
	}
}

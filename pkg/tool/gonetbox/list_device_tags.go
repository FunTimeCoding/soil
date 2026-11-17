package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func listDeviceTags(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "list-device-tags [device]",
		Short: "List the tag names on a device",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(c.ListDeviceTags(context.Background(), arguments[0]))
		},
	}
}

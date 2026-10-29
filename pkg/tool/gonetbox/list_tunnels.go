package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func listTunnels(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "list-tunnels",
		Short: "List all NetBox tunnels",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.EmitValue(c.ListTunnels(context.Background()))
		},
	}
}

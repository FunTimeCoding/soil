package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listTunnels(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "list-tunnels",
		Short: "List all NetBox tunnels",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.Emit(c.ListTunnels())
		},
	}
}

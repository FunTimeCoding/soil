package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listTunnelTerminations(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "list-tunnel-terminations",
		Short: "List all NetBox tunnel terminations",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.Emit(c.ListTunnelTerminations())
		},
	}
}

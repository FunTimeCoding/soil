package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listTunnelGroups(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "list-tunnel-groups",
		Short: "List all NetBox tunnel groups",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.Emit(c.ListTunnelGroups())
		},
	}
}

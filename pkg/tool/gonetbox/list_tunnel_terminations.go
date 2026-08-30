package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listTunnelTerminations(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-tunnel-terminations",
		Short: "List all NetBox tunnel terminations",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.ListTunnelTerminations())
		},
	}
}

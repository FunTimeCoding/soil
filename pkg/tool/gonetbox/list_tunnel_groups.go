package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listTunnelGroups(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-tunnel-groups",
		Short: "List all NetBox tunnel groups",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.ListTunnelGroups())
		},
	}
}

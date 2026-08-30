package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listTunnels(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-tunnels",
		Short: "List all NetBox tunnels",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.ListTunnels())
		},
	}
}

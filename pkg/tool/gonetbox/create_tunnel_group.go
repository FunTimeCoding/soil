package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createTunnelGroup(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "create-tunnel-group [name]",
		Short: "Create a NetBox tunnel group",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(c.CreateTunnelGroup(arguments[0]))
		},
	}
}

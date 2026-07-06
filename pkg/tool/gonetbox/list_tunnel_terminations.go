package gonetbox

import (
	"fmt"
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
			fmt.Println(c.ListTunnelTerminations())
		},
	}
}

package gonetbox

import (
	"fmt"
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
			fmt.Println(c.ListTunnelGroups())
		},
	}
}

package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listClusters(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-clusters",
		Short: "List all NetBox clusters",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.ListClusters())
		},
	}
}

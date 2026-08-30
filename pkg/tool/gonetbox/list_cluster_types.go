package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listClusterTypes(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-cluster-types",
		Short: "List all NetBox cluster types",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.ListClusterTypes())
		},
	}
}

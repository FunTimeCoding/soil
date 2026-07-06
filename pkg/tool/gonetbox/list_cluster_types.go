package gonetbox

import (
	"fmt"
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
			fmt.Println(c.ListClusterTypes())
		},
	}
}

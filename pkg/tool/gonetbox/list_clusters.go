package gonetbox

import (
	"fmt"
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
			fmt.Println(c.ListClusters())
		},
	}
}

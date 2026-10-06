package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listClusters(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "list-clusters",
		Short: "List all NetBox clusters",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.Emit(c.ListClusters())
		},
	}
}

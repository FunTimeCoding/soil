package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listClusterTypes(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "list-cluster-types",
		Short: "List all NetBox cluster types",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.Emit(c.ListClusterTypes())
		},
	}
}

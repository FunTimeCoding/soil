package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listPrefixes(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "list-prefixes",
		Short: "List all NetBox IP prefixes",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.Emit(c.ListPrefixes())
		},
	}
}

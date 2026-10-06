package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listCables(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "list-cables",
		Short: "List all NetBox cables",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.Emit(c.ListCables())
		},
	}
}

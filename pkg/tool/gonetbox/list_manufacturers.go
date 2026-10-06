package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listManufacturers(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "list-manufacturers",
		Short: "List all NetBox manufacturers",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.Emit(c.ListManufacturers())
		},
	}
}

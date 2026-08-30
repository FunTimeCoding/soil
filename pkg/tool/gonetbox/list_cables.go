package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listCables(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-cables",
		Short: "List all NetBox cables",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.ListCables())
		},
	}
}

package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listPrefixes(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-prefixes",
		Short: "List all NetBox IP prefixes",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.ListPrefixes())
		},
	}
}

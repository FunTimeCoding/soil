package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listTags(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-tags",
		Short: "List all NetBox tags",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.ListTags())
		},
	}
}

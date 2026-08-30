package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listPlatforms(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-platforms",
		Short: "List all platforms",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.ListPlatforms())
		},
	}
}

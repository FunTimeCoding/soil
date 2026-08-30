package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createTag(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "create-tag [name]",
		Short: "Create a NetBox tag",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(c.CreateTag(arguments[0]))
		},
	}
}

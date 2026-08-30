package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createSite(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "create-site [name]",
		Short: "Create a NetBox site",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(c.CreateSite(arguments[0]))
		},
	}
}

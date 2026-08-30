package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createManufacturer(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "create-manufacturer [name]",
		Short: "Create a NetBox manufacturer",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(c.CreateManufacturer(arguments[0]))
		},
	}
}

package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listManufacturers(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-manufacturers",
		Short: "List all NetBox manufacturers",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.ListManufacturers())
		},
	}
}

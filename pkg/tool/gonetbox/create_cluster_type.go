package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createClusterType(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "create-cluster-type [name]",
		Short: "Create a NetBox cluster type",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(c.CreateClusterType(arguments[0]))
		},
	}
}

package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createTenant(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "create-tenant [name]",
		Short: "Create a NetBox tenant",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.CreateTenant(arguments[0]))
		},
	}
}

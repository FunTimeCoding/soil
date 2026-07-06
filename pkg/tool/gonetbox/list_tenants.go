package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listTenants(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-tenants",
		Short: "List all NetBox tenants",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(c.ListTenants())
		},
	}
}

package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listTenants(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "list-tenants",
		Short: "List all NetBox tenants",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.Emit(c.ListTenants())
		},
	}
}

package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listCustomFields(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "list-custom-fields",
		Short: "List custom fields",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.Emit(c.ListCustomFields())
		},
	}
}

package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listConfigurationContexts(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "list-configuration-contexts",
		Short: "List configuration contexts with their payload and assignment",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.Emit(c.ListConfigurationContexts())
		},
	}
}

package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func deleteConfigurationContext(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-configuration-context [name]",
		Short: "Delete a configuration context",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.Emit(c.DeleteConfigurationContext(arguments[0]))
		},
	}
}

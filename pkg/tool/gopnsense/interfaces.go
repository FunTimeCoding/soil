package gopnsense

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/client"
	"github.com/spf13/cobra"
)

func interfaces(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "interfaces",
		Short: "List network interfaces",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.Emit(c.Interfaces())
		},
	}
}

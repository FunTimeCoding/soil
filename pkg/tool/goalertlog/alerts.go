package goalertlog

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/client"
	"github.com/spf13/cobra"
)

func alerts(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "alerts <name>",
		Short: "List every recorded alert with this name",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			t.Emit(c.Alerts(a[0]))
		},
	}
}

package goalertlog

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/client"
	"github.com/spf13/cobra"
)

func alerts(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "alerts",
		Short: "List recent alerts",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.Alerts())
		},
	}
}

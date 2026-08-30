package gomaintlog

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/client"
	"github.com/spf13/cobra"
)

func entries(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "entries",
		Short: "List maintenance log entries",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.Entries())
		},
	}
}

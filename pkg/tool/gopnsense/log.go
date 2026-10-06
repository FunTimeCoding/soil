package gopnsense

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/client"
	"github.com/spf13/cobra"
)

func log(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var limit int
	result := &cobra.Command{
		Use:   "log",
		Short: "Read recent firewall log records",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			var l *int

			if limit > 0 {
				l = &limit
			}

			t.Emit(c.Log(l))
		},
	}
	result.Flags().IntVar(&limit, "limit", 0, "maximum number of records")

	return result
}

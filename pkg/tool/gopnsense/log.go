package gopnsense

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/client"
	"github.com/spf13/cobra"
)

func log(c *client.Client) *cobra.Command {
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

			console.Emit(c.Log(l))
		},
	}
	result.Flags().IntVar(&limit, "limit", 0, "maximum number of records")

	return result
}

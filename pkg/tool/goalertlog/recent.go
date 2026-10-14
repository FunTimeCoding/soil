package goalertlog

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	soilTime "github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/client"
	"github.com/spf13/cobra"
	"time"
)

func recent(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var start string
	var end string
	result := &cobra.Command{
		Use:   "recent",
		Short: "List alerts that started in a window, the last hour by default",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			var from *time.Time
			var to *time.Time

			if start != "" {
				from = new(soilTime.Parse(time.RFC3339, start))
			}

			if end != "" {
				to = new(soilTime.Parse(time.RFC3339, end))
			}

			t.Emit(c.RecentAlerts(from, to))
		},
	}
	result.Flags().StringVar(&start, "start", "", "window start, RFC3339")
	result.Flags().StringVar(&end, "end", "", "window end, RFC3339")

	return result
}

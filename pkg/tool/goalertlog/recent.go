package goalertlog

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	soilTime "github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/spf13/cobra"
	"time"
)

func recent(
	c face.Client,
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
			p := &client.GetRecentAlertsParams{}

			if start != "" {
				p.Start = new(soilTime.Parse(time.RFC3339, start))
			}

			if end != "" {
				p.End = new(soilTime.Parse(time.RFC3339, end))
			}

			t.EmitValue(c.GetRecentAlerts(context.Background(), p))
		},
	}
	result.Flags().StringVar(&start, "start", "", "window start, RFC3339")
	result.Flags().StringVar(&end, "end", "", "window end, RFC3339")

	return result
}

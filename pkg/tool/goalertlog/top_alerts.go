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

func topAlerts(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var count int
	var start string
	var end string
	result := &cobra.Command{
		Use:   "top-alerts",
		Short: "List the most frequent alerts in a window",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			to := time.Now()

			if end != "" {
				to = soilTime.Parse(time.RFC3339, end)
			}

			from := to.AddDate(0, 0, -1)

			if start != "" {
				from = soilTime.Parse(time.RFC3339, start)
			}

			t.EmitValue(
				c.GetTopAlerts(
					context.Background(),
					&client.GetTopAlertsParams{
						N:     &count,
						Start: &from,
						End:   &to,
					},
				),
			)
		},
	}
	result.Flags().IntVar(&count, "count", 10, "how many alerts to return")
	result.Flags().StringVar(
		&start,
		"start",
		"",
		"window start, RFC3339 (default one day before end)",
	)
	result.Flags().StringVar(
		&end,
		"end",
		"",
		"window end, RFC3339 (default now)",
	)

	return result
}

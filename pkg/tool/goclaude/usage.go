package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	library "github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/usage_result"
	"github.com/spf13/cobra"
	"time"
)

func usage(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "usage",
		Short: "Show current Claude usage",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			j, found, e := c.Client().GetUsage(context.Background())

			if e != nil {
				c.Terminal().Exitln(e)

				return
			}

			if !found {
				console.Line("No usage data yet.")

				return
			}

			var fableResetAt time.Time

			if j.FableResetAt != nil {
				fableResetAt = *j.FableResetAt
			}

			r := usage_result.New(
				j.FiveHourPercent,
				j.FiveHourReset,
				j.SevenDayPercent,
				j.SevenDayReset,
				j.FablePercent,
				j.FableReset,
				fableResetAt,
				j.LastUpdated,
			)

			if r.HasSession() {
				console.Format(
					"Session  %2d%%   resets %s\n",
					r.FiveHourPercent,
					r.FiveHourResetText(),
				)
			}

			console.Format(
				"Weekly   %2d%%   resets %s\n",
				r.SevenDayPercent,
				r.SevenDayResetText(),
			)

			if r.HasFable() {
				console.Format(
					"Fable    %2d%%   resets %s\n",
					r.FablePercent,
					r.FableResetText(),
				)
			}

			console.Format(
				"Updated  %s\n",
				library.FormatCompact(r.LastUpdated),
			)
		},
	}
}

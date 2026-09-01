package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/usage_result"
	"github.com/spf13/cobra"
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
			response, e := c.Client().GetUsageWithResponse(context.Background())
			errors.PanicOnError(e)

			if response.StatusCode() == 204 || response.JSON200 == nil {
				console.Line("No usage data yet.")

				return
			}

			j := response.JSON200
			r := usage_result.New(
				j.FiveHourPercent,
				j.FiveHourReset,
				j.SevenDayPercent,
				j.SevenDayReset,
				j.FablePercent,
				j.FableReset,
				j.LastUpdated,
			)
			console.Format(
				"Session  %2d%%   resets %s\n",
				r.FiveHourPercent,
				r.FiveHourResetText(),
			)
			console.Format(
				"Weekly   %2d%%   resets %s\n",
				r.SevenDayPercent,
				r.SevenDayResetText(),
			)

			if r.HasFable() {
				console.Format(
					"Fable    %2d%%   resets %s\n",
					r.FablePercent,
					r.FableReset,
				)
			}

			console.Format("Updated  %s\n", time.FormatCompact(r.LastUpdated))
		},
	}
}

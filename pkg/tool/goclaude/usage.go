package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/spf13/cobra"
)

func usage(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "usage",
		Short: "Show current Claude usage from browser monitoring",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			response, e := c.Client().GetUsageWithResponse(
				context.Background(),
			)
			errors.PanicOnError(e)

			if response.StatusCode() == 204 || response.JSON200 == nil {
				fmt.Println("Usage monitoring not enabled or no data yet.")

				return
			}

			r := response.JSON200
			fmt.Printf(
				"Session    %2d%%   resets %s\n",
				r.SessionPercent,
				r.SessionReset,
			)
			fmt.Printf(
				"Weekly     %2d%%   resets %s\n",
				r.WeeklyAllPercent,
				r.WeeklyAllReset,
			)
			fmt.Printf("  Sonnet   %2d%%\n", r.WeeklySonnetPercent)
			fmt.Printf("  Design   %2d%%\n", r.WeeklyDesignPercent)
			fmt.Printf("Routines   %s\n", r.RoutineRuns)
			fmt.Printf(
				"Credits    %s   resets %s\n",
				r.CreditSpent,
				r.CreditReset,
			)
		},
	}
}

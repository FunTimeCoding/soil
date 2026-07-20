package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/pricing"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/spf13/cobra"
)

func cost(c *command_context.Context) *cobra.Command {
	var days int
	result := &cobra.Command{
		Use:   "cost",
		Short: "Show token cost across all sessions for the trailing window",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			response, e := c.Client().GetCostWithResponse(
				context.Background(),
				&client.GetCostParams{Days: &days},
			)
			errors.PanicOnError(e)

			if response.JSON200 == nil {
				return
			}

			r := response.JSON200

			if r.Models == nil || len(*r.Models) == 0 {
				fmt.Printf("No token usage in the last %d day(s).\n", days)

				return
			}

			for _, u := range *r.Models {
				marker := ""

				if !pricing.KnownModel(u.Model) {
					marker = " (unknown model, sonnet rates)"
				}

				fmt.Printf(
					"  %-8s %d calls, %s input, %s output, %s cache-write, %s cache-read, $%.2f%s\n",
					u.Model,
					u.Calls,
					pricing.FormatTokens(u.Input),
					pricing.FormatTokens(u.Output),
					pricing.FormatTokens(u.CacheCreation),
					pricing.FormatTokens(u.CacheRead),
					u.Cost,
					marker,
				)
			}

			fmt.Printf("\nTotal: $%.2f over %d day(s)\n", r.Cost, days)
		},
	}
	result.Flags().IntVar(&days, "days", 1, "trailing window in days")

	return result
}

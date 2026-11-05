package goclaude

import (
	"context"
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
			c.Terminal().EmitValue(
				c.Client().GetCost(
					context.Background(),
					&client.GetCostParams{Days: &days},
				),
			)
		},
	}
	result.Flags().IntVar(&days, "days", 1, "trailing window in days")

	return result
}

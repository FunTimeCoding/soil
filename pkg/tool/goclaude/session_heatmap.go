package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/spf13/cobra"
)

func sessionHeatmap(c *command_context.Context) *cobra.Command {
	var bash bool
	result := &cobra.Command{
		Use:   "heatmap",
		Short: "Aggregate tool usage across sessions",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			c.Terminal().EmitValue(
				c.Client().GetSessionsHeatmap(
					context.Background(),
					&client.GetSessionsHeatmapParams{Bash: &bash},
				),
			)
		},
	}
	result.Flags().BoolVar(
		&bash,
		"bash",
		false,
		"Break down Bash calls by command",
	)

	return result
}

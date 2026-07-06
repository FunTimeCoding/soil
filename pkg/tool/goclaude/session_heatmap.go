package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
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
			response, e := c.Client().GetSessionsHeatmapWithResponse(
				context.Background(),
				&client.GetSessionsHeatmapParams{
					Bash: &bash,
				},
			)
			errors.PanicOnError(e)
			h := response.JSON200

			if len(h.Entries) == 0 {
				fmt.Println("no tool calls found")

				return
			}

			label := "tools"

			if bash {
				label = "commands"
			}

			fmt.Printf(
				"%d sessions (%d with tool calls), %d total calls, %d distinct %s\n\n",
				h.SessionCount,
				h.SessionsWithCalls,
				h.TotalCalls,
				len(h.Entries),
				label,
			)
			fmt.Printf(
				"%-40s %6s %8s\n",
				"Tool",
				"Calls",
				"Sessions",
			)
			fmt.Printf(
				"%-40s %6s %8s\n",
				"----",
				"-----",
				"--------",
			)

			for _, entry := range h.Entries {
				fmt.Printf(
					"%-40s %6d %8d\n",
					entry.Name,
					entry.Calls,
					entry.Sessions,
				)
			}
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

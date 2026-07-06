package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/spf13/cobra"
)

func sessionTools(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "tools <id-or-name>",
		Short: "Show tool usage counts for a session",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier := resolveSession(c.Client(), arguments[0])

			if identifier == "" {
				fmt.Printf(
					"session not found: %s\n",
					arguments[0],
				)

				return
			}

			response, e := c.Client().GetSessionToolsWithResponse(
				context.Background(),
				identifier,
			)
			errors.PanicOnError(e)
			t := response.JSON200

			if len(t.Counts) == 0 {
				fmt.Println("no tool calls found")

				return
			}

			fmt.Printf("%-40s %6s\n", "Tool", "Calls")
			fmt.Printf("%-40s %6s\n", "----", "-----")

			for _, entry := range t.Counts {
				fmt.Printf(
					"%-40s %6d\n",
					entry.Name,
					entry.Count,
				)
			}

			fmt.Printf(
				"%-40s %6d  (%d distinct)\n",
				"TOTAL",
				t.Total,
				len(t.Counts),
			)
		},
	}
}

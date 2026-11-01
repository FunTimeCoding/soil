package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
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
			identifier := resolveSession(c, arguments[0])

			if identifier == "" {
				console.Format("session not found: %s\n", arguments[0])

				return
			}

			t, e := c.Client().GetSessionTools(context.Background(), identifier)

			if e != nil {
				c.Terminal().Exitln(e)

				return
			}

			if len(t.Counts) == 0 {
				console.Line("no tool calls found")

				return
			}

			console.Format("%-40s %6s\n", "Tool", "Calls")
			console.Format("%-40s %6s\n", "----", "-----")

			for _, entry := range t.Counts {
				console.Format("%-40s %6d\n", entry.Name, entry.Count)
			}

			console.Format(
				"%-40s %6d  (%d distinct)\n",
				"TOTAL",
				t.Total,
				len(t.Counts),
			)
		},
	}
}

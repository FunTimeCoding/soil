package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/spf13/cobra"
	"strings"
)

func sessionContext(c *command_context.Context) *cobra.Command {
	var surround int
	result := &cobra.Command{
		Use:   "context <id-or-name> <tool-filter>",
		Short: "Show conversation context around tool uses",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier := resolveSession(c, arguments[0])

			if identifier == "" {
				return
			}

			response, e := c.Client().GetSessionToolContext(
				context.Background(),
				identifier,
				&client.GetSessionToolContextParams{
					Filter:   arguments[1],
					Surround: &surround,
				},
			)

			if e != nil {
				c.Terminal().Exitln(e)

				return
			}

			t := c.Terminal()
			results := response.Results

			if len(results) == 0 {
				t.Line(fmt.Sprintf("no tool calls matching %q", arguments[1]))

				return
			}

			for i, r := range results {
				if i > 0 {
					t.Line(strings.Repeat("─", 60))
				}

				t.Line(fmt.Sprintf("### %s", r.ToolName))
				t.Line()

				if r.Before != nil {
					for _, m := range *r.Before {
						printContextMessage(t, m)
					}
				}

				t.Line(fmt.Sprintf("**→ %s called**", r.ToolName))
				t.Line()

				if r.After != nil {
					for _, m := range *r.After {
						printContextMessage(t, m)
					}
				}
			}
		},
	}
	result.Flags().IntVar(
		&surround,
		"surround",
		1,
		"Number of messages before and after each tool use",
	)

	return result
}

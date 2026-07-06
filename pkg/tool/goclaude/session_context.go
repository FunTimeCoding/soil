package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
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
			identifier := resolveSession(c.Client(), arguments[0])

			if identifier == "" {
				fmt.Printf(
					"session not found: %s\n",
					arguments[0],
				)

				return
			}

			response, e := c.Client().GetSessionToolContextWithResponse(
				context.Background(),
				identifier,
				&client.GetSessionToolContextParams{
					Filter:   arguments[1],
					Surround: &surround,
				},
			)
			errors.PanicOnError(e)
			results := response.JSON200.Results

			if len(results) == 0 {
				fmt.Printf(
					"no tool calls matching %q\n",
					arguments[1],
				)

				return
			}

			for i, r := range results {
				if i > 0 {
					fmt.Println(strings.Repeat("─", 60))
				}

				fmt.Printf("### %s\n\n", r.ToolName)

				if r.Before != nil {
					for _, m := range *r.Before {
						text := m.Text

						if len(text) > 300 {
							text = fmt.Sprintf(
								"%s…",
								text[:300],
							)
						}

						fmt.Printf(
							"**%s:** %s\n\n",
							m.Role,
							text,
						)
					}
				}

				fmt.Printf("**→ %s called**\n\n", r.ToolName)

				if r.After != nil {
					for _, m := range *r.After {
						text := m.Text

						if len(text) > 300 {
							text = fmt.Sprintf(
								"%s…",
								text[:300],
							)
						}

						fmt.Printf(
							"**%s:** %s\n\n",
							m.Role,
							text,
						)
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

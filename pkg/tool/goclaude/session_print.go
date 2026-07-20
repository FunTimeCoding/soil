package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/spf13/cobra"
)

func sessionPrint(c *command_context.Context) *cobra.Command {
	var tail int
	var turns string
	result := &cobra.Command{
		Use:   "print <id-or-name>",
		Short: "Print a session to stdout",
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

			response, e := c.Client().GetSessionMessagesWithResponse(
				context.Background(),
				identifier,
			)
			errors.PanicOnError(e)
			var messages []client.SessionMessage
			var turnNumbers []int
			turn := 0

			for _, m := range response.JSON200.Messages {
				if m.Text == "" {
					continue
				}

				if m.IsMeta != nil && *m.IsMeta {
					continue
				}

				if m.Role == "user" {
					turn++
				}

				messages = append(messages, m)
				turnNumbers = append(turnNumbers, turn)
			}

			from := 0
			to := 0

			if turns != "" {
				okay := false
				from, to, okay = parseTurnRange(turns)

				if !okay {
					fmt.Printf("invalid --turns range: %s\n", turns)

					return
				}
			}

			if tail > 0 {
				from = turn - tail + 1

				if from < 1 {
					from = 1
				}

				to = turn
			}

			for i, m := range messages {
				if from > 0 && (turnNumbers[i] < from || turnNumbers[i] > to) {
					continue
				}

				if m.Role == "user" {
					fmt.Printf(
						"## user [%d]\n\n%s\n\n",
						turnNumbers[i],
						m.Text,
					)

					continue
				}

				fmt.Printf("## %s\n\n%s\n\n", m.Role, m.Text)
			}
		},
	}
	result.Flags().IntVar(
		&tail,
		"tail",
		0,
		"Only the last N user turns with their full replies",
	)
	result.Flags().StringVar(
		&turns,
		"turns",
		"",
		"Only the given user turn range, e.g. 57 or 55-60",
	)
	result.MarkFlagsMutuallyExclusive("tail", "turns")

	return result
}

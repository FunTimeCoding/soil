package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/spf13/cobra"
)

func sessionPrint(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
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

			for _, m := range response.JSON200.Messages {
				if m.Text == "" {
					continue
				}

				if m.IsMeta != nil && *m.IsMeta {
					continue
				}

				fmt.Printf("## %s\n\n%s\n\n", m.Role, m.Text)
			}
		},
	}
}

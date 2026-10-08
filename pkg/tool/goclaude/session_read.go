package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/spf13/cobra"
)

func sessionRead(c *command_context.Context) *cobra.Command {
	var count int
	result := &cobra.Command{
		Use:   "read <id-or-name> <block>",
		Short: "Read the blocks around one search hit",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier := resolveSession(c.Client(), arguments[0])

			if identifier == "" {
				console.Format("session not found: %s\n", arguments[0])

				return
			}

			parameters := &client.GetSessionWindowParams{Around: arguments[1]}

			if count > 0 {
				parameters.Count = &count
			}

			response, e := c.Client().GetSessionWindowWithResponse(
				context.Background(),
				identifier,
				parameters,
			)
			errors.PanicOnError(e)

			if response.JSON404 != nil {
				console.Format("%s\n", response.JSON404.Error)

				return
			}

			if response.JSON200 == nil {
				c.Terminal().Reject(response.Status(), response.Body)
			}

			for _, b := range response.JSON200.Blocks {
				marker := ""

				if b.Identifier == arguments[1] {
					marker = " <- hit"
				}

				console.Format(
					"[%s %s %s %s]%s\n%s\n\n",
					b.At,
					b.Role,
					b.Kind,
					b.Identifier,
					marker,
					b.Text,
				)
			}
		},
	}
	result.Flags().IntVar(
		&count,
		"count",
		0,
		"Blocks on each side of the hit (default 5, at most 25)",
	)

	return result
}

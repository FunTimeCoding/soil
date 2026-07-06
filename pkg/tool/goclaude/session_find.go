package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/spf13/cobra"
)

func sessionFind(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "find <tool-filter>",
		Short: "Find sessions by tool usage",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			response, e := c.Client().GetSessionsFindWithResponse(
				context.Background(),
				&client.GetSessionsFindParams{
					Tool: arguments[0],
				},
			)
			errors.PanicOnError(e)
			matches := response.JSON200.Matches

			if len(matches) == 0 {
				fmt.Printf(
					"no sessions with tool calls matching %q\n",
					arguments[0],
				)

				return
			}

			for _, m := range matches {
				fmt.Printf(
					"%4d  %.8s  %s\n",
					m.Count,
					m.SessionId,
					m.Name,
				)
			}
		},
	}
}

package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/spf13/cobra"
)

func sessionShow(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "show <id-or-name>",
		Short: "Show session detail with alias, description, completions, and summary",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier := resolveSession(c, arguments[0])

			if identifier == "" {
				return
			}

			c.Terminal().EmitValue(
				c.Client().GetSessionDetail(context.Background(), identifier),
			)
		},
	}
}

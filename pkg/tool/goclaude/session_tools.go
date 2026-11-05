package goclaude

import (
	"context"
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
				return
			}

			c.Terminal().EmitValue(
				c.Client().GetSessionTools(context.Background(), identifier),
			)
		},
	}
}

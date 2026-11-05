package goclaude

import (
	"context"
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
			c.Terminal().EmitValue(
				c.Client().GetSessionsFind(
					context.Background(),
					&client.GetSessionsFindParams{Tool: arguments[0]},
				),
			)
		},
	}
}

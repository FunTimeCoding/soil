package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/spf13/cobra"
)

func sessionList(c *command_context.Context) *cobra.Command {
	var peek bool
	var unnamed bool
	var limit int
	var offset int
	result := &cobra.Command{
		Use:   "list",
		Short: "List all sessions",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			params := &client.GetSessionsParams{Peek: &peek}

			if unnamed {
				params.Unnamed = &unnamed
			}

			if limit > 0 {
				params.Limit = &limit
			}

			if offset > 0 {
				params.Offset = &offset
			}

			c.Terminal().EmitValue(
				c.Client().GetSessions(context.Background(), params),
			)
		},
	}
	result.Flags().BoolVar(
		&peek,
		"peek",
		false,
		"Show first user message preview",
	)
	result.Flags().BoolVar(
		&unnamed,
		"unnamed",
		false,
		"Only sessions without an alias, oldest first",
	)
	result.Flags().IntVar(
		&limit,
		"limit",
		0,
		"Maximum number of sessions to show",
	)
	result.Flags().IntVar(&offset, "offset", 0, "Number of sessions to skip")

	return result
}

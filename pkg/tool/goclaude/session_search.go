package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/spf13/cobra"
	"strings"
)

func sessionSearch(c *command_context.Context) *cobra.Command {
	var kinds []string
	var limit int
	result := &cobra.Command{
		Use:   "search <term>...",
		Short: "Find conversations holding every term",
		Args:  cobra.MinimumNArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			parameters := &client.GetSessionsSearchParams{
				Query: strings.Join(arguments, " "),
			}

			if len(kinds) > 0 {
				parameters.Kinds = &kinds
			}

			if limit > 0 {
				parameters.Limit = &limit
			}

			c.Terminal().EmitValue(
				c.Client().GetSessionsSearch(context.Background(), parameters),
			)
		},
	}
	result.Flags().StringSliceVar(
		&kinds,
		"kind",
		nil,
		"Block kinds to search: message, edit, call (default message)",
	)
	result.Flags().IntVar(
		&limit,
		"limit",
		0,
		"Maximum conversations (default 20, at most 100)",
	)

	return result
}

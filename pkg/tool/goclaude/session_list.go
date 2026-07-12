package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/spf13/cobra"
)

func sessionList(c *command_context.Context) *cobra.Command {
	var peek bool
	var detail bool
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

			response, e := c.Client().GetSessionsWithResponse(
				context.Background(),
				params,
			)
			errors.PanicOnError(e)

			for _, s := range response.JSON200.Sessions {
				ts := formatTimestamp(s.Timestamp)
				alias := ""

				if s.Alias != nil {
					alias = *s.Alias
				}

				slug := ""

				if s.Slug != nil {
					slug = *s.Slug
				}

				poolName := ""

				if s.Name != nil {
					poolName = *s.Name
				}

				displayName := slug
				marker := " "

				if alias != "" {
					displayName = alias
					marker = "*"
				}

				shortIdentifier := s.Identifier

				if len(shortIdentifier) > 8 {
					shortIdentifier = shortIdentifier[:8]
				}

				fmt.Printf(
					"%s  %s  %-7s %5d  %s %s",
					ts,
					shortIdentifier,
					poolName,
					s.Lines,
					marker,
					displayName,
				)

				if peek && s.Preview != nil {
					fmt.Printf("  %s", *s.Preview)
				}

				if detail && s.Description != nil && *s.Description != "" {
					fmt.Printf("\n    %s", *s.Description)
				}

				fmt.Println()
			}
		},
	}
	result.Flags().BoolVar(
		&peek,
		"peek",
		false,
		"Show first user message preview",
	)
	result.Flags().BoolVar(
		&detail,
		"detail",
		false,
		"Show descriptions",
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
	result.Flags().IntVar(
		&offset,
		"offset",
		0,
		"Number of sessions to skip",
	)

	return result
}

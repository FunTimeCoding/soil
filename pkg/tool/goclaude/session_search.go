package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
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

			response, e := c.Client().GetSessionsSearchWithResponse(
				context.Background(),
				parameters,
			)
			errors.PanicOnError(e)

			if response.JSON200 == nil {
				c.Terminal().Reject(response.Status(), response.Body)
			}

			found := response.JSON200

			if found.Indexed < found.Total {
				console.Format(
					"indexing %d/%d conversations - results may be incomplete\n",
					found.Indexed,
					found.Total,
				)
			}

			if len(found.Conversations) == 0 {
				console.Format("no conversation holds every term\n")

				return
			}

			for _, v := range found.Conversations {
				name := v.Name

				if name == "" {
					name = constant.UnnamedSession
				}

				console.Format(
					"%4d  %.8s  %s  %s\n",
					v.Count,
					v.Session,
					name,
					v.Latest,
				)

				for _, h := range v.Hits {
					console.Format(
						"      %s %-9s %-7s %s\n        %s\n",
						h.At,
						h.Role,
						h.Kind,
						h.Identifier,
						h.Snippet,
					)
				}
			}
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

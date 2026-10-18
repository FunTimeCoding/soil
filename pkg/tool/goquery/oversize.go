package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"github.com/spf13/cobra"
)

func oversize(
	c *client.ClientWithResponses,
	t *terminal.Terminal,
) *cobra.Command {
	var collection string
	result := &cobra.Command{
		Use:   "oversize",
		Short: "List chunks too big for the reranker window, exit 1 if any",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			params := &client.GetOversizeParams{}

			if collection != "" {
				params.Collection = &collection
			}

			r, e := c.GetOversizeWithResponse(context.Background(), params)
			errors.PanicOnError(e)
			refusal.PanicOnFailure("oversize", r.HTTPResponse, r.Body)
			printOversize(r.JSON200)

			if len(r.JSON200.Files) > 0 {
				t.Exit(1)
			}
		},
	}
	result.Flags().StringVar(
		&collection,
		"collection",
		"",
		"Restrict the check to a collection",
	)

	return result
}

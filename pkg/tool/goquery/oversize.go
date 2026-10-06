package goquery

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/spf13/cobra"
)

func oversize(
	c *client.Client,
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

			r, e := c.GetOversize(context.Background(), params)
			errors.PanicOnError(e)
			p, f := client.ParseGetOversizeResponse(r)
			errors.PanicOnError(f)

			if p.JSON200 == nil {
				panic(fmt.Sprintf("oversize: %s: %s", p.Status(), p.Body))
			}

			printOversize(p.JSON200)

			if len(p.JSON200.Files) > 0 {
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

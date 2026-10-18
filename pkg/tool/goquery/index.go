package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"github.com/spf13/cobra"
)

func index(c *client.ClientWithResponses) *cobra.Command {
	var collection string
	result := &cobra.Command{
		Use:   "index",
		Short: "Re-index collections",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			body := client.PostIndexJSONRequestBody{}

			if collection != "" {
				body.Collection = &collection
			}

			r, e := c.PostIndexWithResponse(context.Background(), body)
			errors.PanicOnError(e)
			refusal.PanicOnFailure("index", r.HTTPResponse, r.Body)

			for _, v := range *r.JSON200 {
				console.Format(
					"%s: indexed=%d updated=%d unchanged=%d removed=%d\n",
					v.Collection,
					v.Indexed,
					v.Updated,
					v.Unchanged,
					v.Removed,
				)
			}
		},
	}
	result.Flags().StringVar(
		&collection,
		"collection",
		"",
		"Re-index a specific collection only",
	)

	return result
}

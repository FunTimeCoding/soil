package goquery

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/spf13/cobra"
)

func index(c *client.Client) *cobra.Command {
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

			r, e := c.PostIndex(context.Background(), body)
			errors.PanicOnError(e)
			defer errors.PanicClose(r.Body)
			var results []client.IndexResult
			errors.PanicOnError(json.NewDecoder(r.Body).Decode(&results))

			for _, v := range results {
				fmt.Printf(
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

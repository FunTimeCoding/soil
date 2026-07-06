package goquery

import (
	"context"
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/spf13/cobra"
)

func listMetadata(c *client.Client) *cobra.Command {
	var key string
	result := &cobra.Command{
		Use:   "list-metadata [collection]",
		Short: "List metadata keys and values for a collection",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			params := &client.GetMetadataParams{
				Collection: arguments[0],
			}

			if key != "" {
				params.Key = &key
			}

			r, e := c.GetMetadata(context.Background(), params)
			errors.PanicOnError(e)
			defer errors.PanicClose(r.Body)
			var facets []client.Facet
			errors.PanicOnError(json.NewDecoder(r.Body).Decode(&facets))
			printFacets(&facets)
		},
	}
	result.Flags().StringVar(
		&key,
		"key",
		"",
		"Show all values for a specific key",
	)

	return result
}

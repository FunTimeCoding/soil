package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"github.com/spf13/cobra"
)

func listMetadata(c *client.ClientWithResponses) *cobra.Command {
	var key string
	result := &cobra.Command{
		Use:   "list-metadata [collection]",
		Short: "List metadata keys and values for a collection",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			params := &client.GetMetadataParams{Collection: arguments[0]}

			if key != "" {
				params.Key = &key
			}

			r, e := c.GetMetadataWithResponse(context.Background(), params)
			errors.PanicOnError(e)
			refusal.PanicOnFailure("list metadata", r.HTTPResponse, r.Body)
			var facets []client.Facet

			for _, f := range *r.JSON200 {
				facets = append(facets, *f)
			}

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

package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/spf13/cobra"
)

func listMetadata(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
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

			found, e := c.GetMetadata(context.Background(), params)
			t.EmitDone(e)
			var facets []client.Facet

			for _, f := range found {
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

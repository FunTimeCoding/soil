package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"github.com/spf13/cobra"
)

func addCollection(c *client.ClientWithResponses) *cobra.Command {
	var pattern string
	result := &cobra.Command{
		Use:   "add-collection [name] [path]",
		Short: "Add or update a collection",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			body := client.PostCollectionJSONRequestBody{
				Name: arguments[0],
				Path: arguments[1],
			}

			if pattern != "" {
				body.Pattern = &pattern
			}

			r, e := c.PostCollectionWithResponse(context.Background(), body)
			errors.PanicOnError(e)
			refusal.PanicOnFailure("add collection", r.HTTPResponse, r.Body)
			console.Format("collection %s added\n", arguments[0])
		},
	}
	result.Flags().StringVar(
		&pattern,
		"pattern",
		"**/*.md",
		"Glob pattern for files to index",
	)

	return result
}

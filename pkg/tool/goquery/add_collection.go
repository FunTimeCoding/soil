package goquery

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/spf13/cobra"
)

func addCollection(c *client.Client) *cobra.Command {
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

			r, e := c.PostCollection(context.Background(), body)
			errors.PanicOnError(e)
			defer errors.PanicClose(r.Body)
			fmt.Printf("collection %s added\n", arguments[0])
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

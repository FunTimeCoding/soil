package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/spf13/cobra"
)

func deleteCollection(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-collection [name]",
		Short: "Delete a collection and all its documents",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			r, e := c.DeleteCollection(
				context.Background(),
				&client.DeleteCollectionParams{Name: arguments[0]},
			)
			errors.PanicOnError(e)
			defer errors.PanicClose(r.Body)
			console.Format("collection %s deleted\n", arguments[0])
		},
	}
}

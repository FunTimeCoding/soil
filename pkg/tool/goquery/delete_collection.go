package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/spf13/cobra"
)

func deleteCollection(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-collection [name]",
		Short: "Delete a collection and all its documents",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			deleted, e := c.DeleteCollection(
				context.Background(),
				&client.DeleteCollectionParams{Name: arguments[0]},
			)
			t.EmitDone(e)

			if !deleted {
				t.Exitf("collection not found: %s\n", arguments[0])
			}

			console.Format("collection %s deleted\n", arguments[0])
		},
	}
}

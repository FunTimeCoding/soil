package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/spf13/cobra"
)

func addCollection(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
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

			t.EmitDone(c.PostCollection(context.Background(), body))
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

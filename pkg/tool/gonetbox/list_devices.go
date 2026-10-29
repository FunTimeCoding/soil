package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func listDevices(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var query string
	result := &cobra.Command{
		Use:   "list-devices",
		Short: "List NetBox devices",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			var q *string

			if query != "" {
				q = &query
			}

			t.EmitValue(c.ListDevices(context.Background(), q))
		},
	}
	result.Flags().StringVar(&query, "query", "", "filter by name")

	return result
}

package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/spf13/cobra"
)

func removeContext(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "context-remove [collection] [path-prefix]",
		Short: "Remove context for a path prefix",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			deleted, e := c.DeleteContext(
				context.Background(),
				&client.DeleteContextParams{
					Collection: arguments[0],
					PathPrefix: arguments[1],
				},
			)
			t.EmitDone(e)

			if !deleted {
				t.Exitf(
					"context not found: %s %s\n",
					arguments[0],
					arguments[1],
				)
			}

			console.Format(
				"context removed: %s %s\n",
				arguments[0],
				arguments[1],
			)
		},
	}
}

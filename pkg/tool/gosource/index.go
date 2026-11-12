package gosource

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/face"
	"github.com/spf13/cobra"
)

func index(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	result := &cobra.Command{
		Use:   "index",
		Short: "Manage gosourced's source index",
	}
	result.AddCommand(
		&cobra.Command{
			Use:   "prune",
			Short: "Delete every index generation but the daemon's own, and its builds untouched for a day",
			Args:  cobra.NoArgs,
			Run: func(
				_ *cobra.Command,
				_ []string,
			) {
				removed, e := c.PostIndexPrune(context.Background())

				if e != nil {
					t.EmitDone(e)

					return
				}

				for _, l := range PruneLines(removed) {
					t.Line(l)
				}
			},
		},
	)

	return result
}

package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/spf13/cobra"
)

func sessionSweep(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "sweep",
		Short: "Copy session files to safe harbor",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := c.Client().PostSweep(context.Background())

			if e != nil {
				c.Terminal().Exitln(e)

				return
			}

			console.Format(
				"sweep: %d copied, %d updated, %d skipped\n",
				r.Copied,
				r.Updated,
				r.Skipped,
			)
		},
	}
}

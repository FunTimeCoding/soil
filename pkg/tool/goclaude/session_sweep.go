package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
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
			response, e := c.Client().PostSweepWithResponse(
				context.Background(),
			)
			errors.PanicOnError(e)
			r := response.JSON200
			fmt.Printf(
				"sweep: %d copied, %d updated, %d skipped\n",
				r.Copied,
				r.Updated,
				r.Skipped,
			)
		},
	}
}

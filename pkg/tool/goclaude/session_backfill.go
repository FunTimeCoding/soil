package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/spf13/cobra"
)

func sessionBackfill(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "backfill",
		Short: "Re-enrich all sessions from JSONL files",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			response, e := c.Client().PostBackfillWithResponse(
				context.Background(),
			)
			errors.PanicOnError(e)
			r := response.JSON200
			fmt.Printf(
				"backfill: %d enriched, %d skipped\n",
				r.Enriched,
				r.Skipped,
			)
		},
	}
}

package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/spf13/cobra"
)

func sessionBackfill(c *command_context.Context) *cobra.Command {
	var cold bool
	result := &cobra.Command{
		Use:   "backfill",
		Short: "Re-enrich all sessions from JSONL files",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			parameters := &client.PostBackfillParams{}

			if cold {
				parameters.Cold = new(true)
			}

			c.Terminal().EmitValue(
				c.Client().PostBackfill(context.Background(), parameters),
			)
		},
	}
	result.Flags().BoolVar(
		&cold,
		"cold",
		false,
		"Reset tracker offsets and re-read every transcript whole",
	)

	return result
}

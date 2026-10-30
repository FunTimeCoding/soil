package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/spf13/cobra"
)

func failed(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var collection string
	result := &cobra.Command{
		Use:   "failed",
		Short: "List chunks the embedder refused, exit 1 if any",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			params := &client.GetFailedParams{}

			if collection != "" {
				params.Collection = &collection
			}

			report, e := c.GetFailed(context.Background(), params)
			t.EmitDone(e)
			printFailed(report)

			if len(report.Chunks) > 0 {
				t.Exit(1)
			}
		},
	}
	result.Flags().StringVar(
		&collection,
		"collection",
		"",
		"Restrict the list to a collection",
	)

	return result
}

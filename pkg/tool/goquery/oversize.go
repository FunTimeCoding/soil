package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/spf13/cobra"
)

func oversize(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var collection string
	result := &cobra.Command{
		Use:   "oversize",
		Short: "List chunks too big for the reranker window, exit 1 if any",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			params := &client.GetOversizeParams{}

			if collection != "" {
				params.Collection = &collection
			}

			report, e := c.GetOversize(context.Background(), params)
			t.EmitDone(e)
			printOversize(report)

			if len(report.Files) > 0 {
				t.Exit(1)
			}
		},
	}
	result.Flags().StringVar(
		&collection,
		"collection",
		"",
		"Restrict the check to a collection",
	)

	return result
}

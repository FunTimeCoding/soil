package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/spf13/cobra"
)

func index(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var collection string
	result := &cobra.Command{
		Use:   "index",
		Short: "Re-index collections",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			body := client.PostIndexJSONRequestBody{}

			if collection != "" {
				body.Collection = &collection
			}

			t.EmitValue(c.PostIndex(context.Background(), body))
		},
	}
	result.Flags().StringVar(
		&collection,
		"collection",
		"",
		"Re-index a specific collection only",
	)

	return result
}

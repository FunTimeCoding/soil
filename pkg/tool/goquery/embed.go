package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"github.com/spf13/cobra"
)

func embed(c *client.ClientWithResponses) *cobra.Command {
	return &cobra.Command{
		Use:   "embed",
		Short: "Generate embeddings for indexed documents",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := c.PostEmbedWithResponse(context.Background())
			errors.PanicOnError(e)
			refusal.PanicOnFailure("embed", r.HTTPResponse, r.Body)
			console.Format(
				"Embedded %d documents (%d chunks)\n",
				r.JSON200.Documents,
				r.JSON200.Chunks,
			)
		},
	}
}

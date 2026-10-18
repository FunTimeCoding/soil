package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"github.com/spf13/cobra"
)

func status(c *client.ClientWithResponses) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Show index status",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := c.GetStatusWithResponse(context.Background())
			errors.PanicOnError(e)
			refusal.PanicOnFailure("status", r.HTTPResponse, r.Body)
			s := r.JSON200
			console.Format(
				"Documents: %d  Embeddings: %d  Pending: %d\n",
				s.TotalDocuments,
				s.TotalEmbeddings,
				s.PendingEmbeddings,
			)

			for _, v := range s.Collections {
				if v.Path == "" {
					console.Format(
						"  %s: %d documents\n",
						v.Name,
						v.DocumentCount,
					)
				} else {
					console.Format(
						"  %s: %d documents (%s %s)\n",
						v.Name,
						v.DocumentCount,
						v.Path,
						v.Pattern,
					)
				}
			}
		},
	}
}

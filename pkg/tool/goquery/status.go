package goquery

import (
	"context"
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/spf13/cobra"
)

func status(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Show index status",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := c.GetStatus(context.Background())
			errors.PanicOnError(e)
			defer errors.PanicClose(r.Body)
			var s client.Status
			errors.PanicOnError(json.NewDecoder(r.Body).Decode(&s))
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

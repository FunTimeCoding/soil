package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/spf13/cobra"
)

func status(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Show index status",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			s, e := c.GetStatus(context.Background())
			t.EmitDone(e)
			console.Format(
				"Documents: %d  Embeddings: %d  Pending: %d  Failed: %d\n",
				s.TotalDocuments,
				s.TotalEmbeddings,
				s.PendingEmbeddings,
				s.FailedEmbeddings,
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

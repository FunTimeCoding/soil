package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"github.com/spf13/cobra"
)

func rechunk(c *client.ClientWithResponses) *cobra.Command {
	return &cobra.Command{
		Use:   "rechunk",
		Short: "Re-embed documents whose stored chunks no longer match the chunker",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := c.PostRechunkWithResponse(context.Background())
			errors.PanicOnError(e)
			refusal.PanicOnFailure("rechunk", r.HTTPResponse, r.Body)

			for _, d := range r.JSON200.Documents {
				console.Format("  %s\n", d)
			}

			console.Format(
				"%d documents queued for re-embedding\n",
				len(r.JSON200.Documents),
			)
			s, f := c.PostEmbedWithResponse(context.Background())
			errors.PanicOnError(f)
			refusal.PanicOnFailure("embed", s.HTTPResponse, s.Body)
			console.Format(
				"Embedded %d documents (%d chunks)\n",
				s.JSON200.Documents,
				s.JSON200.Chunks,
			)
		},
	}
}

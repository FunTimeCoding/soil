package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/spf13/cobra"
)

func rechunk(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "rechunk",
		Short: "Re-embed documents whose stored chunks no longer match the chunker",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := c.PostRechunk(context.Background())
			t.EmitDone(e)

			for _, d := range r.Documents {
				console.Format("  %s\n", d)
			}

			console.Format(
				"%d documents queued for re-embedding\n",
				len(r.Documents),
			)
			s, f := c.PostEmbed(context.Background())
			t.EmitDone(f)
			console.Format(
				"Embedded %d documents (%d chunks)\n",
				s.Documents,
				s.Chunks,
			)
		},
	}
}

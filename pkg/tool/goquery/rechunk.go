package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goquery/result"
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

			if e != nil {
				t.EmitDone(e)

				return
			}

			s, f := c.PostEmbed(context.Background())

			if f != nil {
				t.EmitDone(f)

				return
			}

			t.EmitValue(result.NewRechunk(r.Documents, s), nil)
		},
	}
}

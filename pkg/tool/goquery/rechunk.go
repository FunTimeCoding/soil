package goquery

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/spf13/cobra"
)

func rechunk(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "rechunk",
		Short: "Re-embed documents whose stored chunks no longer match the chunker",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := c.PostRechunk(context.Background())
			errors.PanicOnError(e)
			p, f := client.ParsePostRechunkResponse(r)
			errors.PanicOnError(f)

			if p.JSON200 == nil {
				panic(fmt.Sprintf("rechunk: %s: %s", p.Status(), p.Body))
			}

			for _, d := range p.JSON200.Documents {
				console.Format("  %s\n", d)
			}

			console.Format(
				"%d documents queued for re-embedding\n",
				len(p.JSON200.Documents),
			)
			s, g := c.PostEmbed(context.Background())
			errors.PanicOnError(g)
			defer errors.PanicClose(s.Body)
			var result client.EmbedResult
			errors.PanicOnError(json.NewDecoder(s.Body).Decode(&result))
			console.Format(
				"Embedded %d documents (%d chunks)\n",
				result.Documents,
				result.Chunks,
			)
		},
	}
}

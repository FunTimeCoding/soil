package goquery

import (
	"context"
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/spf13/cobra"
)

func embed(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "embed",
		Short: "Generate embeddings for indexed documents",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := c.PostEmbed(context.Background())
			errors.PanicOnError(e)
			defer errors.PanicClose(r.Body)
			var result client.EmbedResult
			errors.PanicOnError(json.NewDecoder(r.Body).Decode(&result))
			console.Format(
				"Embedded %d documents (%d chunks)\n",
				result.Documents,
				result.Chunks,
			)
		},
	}
}

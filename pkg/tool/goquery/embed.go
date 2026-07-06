package goquery

import (
	"context"
	"encoding/json"
	"fmt"
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
			fmt.Printf(
				"Embedded %d documents (%d chunks)\n",
				result.Documents,
				result.Chunks,
			)
		},
	}
}

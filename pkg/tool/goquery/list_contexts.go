package goquery

import (
	"context"
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/spf13/cobra"
)

func listContexts(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "context-list",
		Short: "List all contexts",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := c.GetContext(context.Background())
			errors.PanicOnError(e)
			defer errors.PanicClose(r.Body)
			var results []client.ContextEntry
			errors.PanicOnError(json.NewDecoder(r.Body).Decode(&results))

			for _, v := range results {
				console.Format(
					"%s  %s  %s\n",
					v.Collection,
					v.PathPrefix,
					v.Description,
				)
			}
		},
	}
}

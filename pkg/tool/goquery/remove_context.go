package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"github.com/spf13/cobra"
)

func removeContext(c *client.ClientWithResponses) *cobra.Command {
	return &cobra.Command{
		Use:   "context-remove [collection] [path-prefix]",
		Short: "Remove context for a path prefix",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			params := &client.DeleteContextParams{
				Collection: arguments[0],
				PathPrefix: arguments[1],
			}
			r, e := c.DeleteContextWithResponse(context.Background(), params)
			errors.PanicOnError(e)
			refusal.PanicOnFailure("remove context", r.HTTPResponse, r.Body)

			if r.JSON200.Deleted != nil && *r.JSON200.Deleted {
				console.Format(
					"context removed: %s %s\n",
					arguments[0],
					arguments[1],
				)
			} else {
				console.Format(
					"context not found: %s %s\n",
					arguments[0],
					arguments[1],
				)
			}
		},
	}
}

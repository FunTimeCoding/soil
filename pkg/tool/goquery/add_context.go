package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"github.com/spf13/cobra"
)

func addContext(c *client.ClientWithResponses) *cobra.Command {
	return &cobra.Command{
		Use:   "context-add [collection] [path-prefix] [description]",
		Short: "Add context for a path prefix within a collection",
		Args:  cobra.ExactArgs(3),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			r, e := c.PostContextWithResponse(
				context.Background(),
				client.PostContextJSONRequestBody{
					Collection:  arguments[0],
					PathPrefix:  arguments[1],
					Description: arguments[2],
				},
			)
			errors.PanicOnError(e)
			refusal.PanicOnFailure("add context", r.HTTPResponse, r.Body)
			console.Format("context added: %s %s\n", arguments[0], arguments[1])
		},
	}
}

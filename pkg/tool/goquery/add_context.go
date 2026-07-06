package goquery

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/spf13/cobra"
)

func addContext(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "context-add [collection] [path-prefix] [description]",
		Short: "Add context for a path prefix within a collection",
		Args:  cobra.ExactArgs(3),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			r, e := c.PostContext(
				context.Background(),
				client.PostContextJSONRequestBody{
					Collection:  arguments[0],
					PathPrefix:  arguments[1],
					Description: arguments[2],
				},
			)
			errors.PanicOnError(e)
			defer errors.PanicClose(r.Body)
			fmt.Printf("context added: %s %s\n", arguments[0], arguments[1])
		},
	}
}

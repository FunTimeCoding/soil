package goquery

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/spf13/cobra"
)

func removeContext(c *client.Client) *cobra.Command {
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
			r, e := c.DeleteContext(context.Background(), params)
			errors.PanicOnError(e)
			defer errors.PanicClose(r.Body)
			var result map[string]bool
			errors.PanicOnError(json.NewDecoder(r.Body).Decode(&result))

			if result["deleted"] {
				fmt.Printf(
					"context removed: %s %s\n",
					arguments[0],
					arguments[1],
				)
			} else {
				fmt.Printf(
					"context not found: %s %s\n",
					arguments[0],
					arguments[1],
				)
			}
		},
	}
}

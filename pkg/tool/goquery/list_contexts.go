package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"github.com/spf13/cobra"
)

func listContexts(c *client.ClientWithResponses) *cobra.Command {
	return &cobra.Command{
		Use:   "context-list",
		Short: "List all contexts",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := c.GetContextWithResponse(context.Background())
			errors.PanicOnError(e)
			refusal.PanicOnFailure("list contexts", r.HTTPResponse, r.Body)

			for _, v := range *r.JSON200 {
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

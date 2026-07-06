package gopostgres

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"github.com/spf13/cobra"
)

func listInstances(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-instances",
		Short: "List configured PostgreSQL instances",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := c.ListInstances(context.Background())
			errors.PanicOnError(e)
			printResponse(r)
		},
	}
}

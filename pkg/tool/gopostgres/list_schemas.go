package gopostgres

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"github.com/spf13/cobra"
)

func listSchemas(c *client.Client) *cobra.Command {
	var instance string
	result := &cobra.Command{
		Use:   "list-schemas",
		Short: "List schemas in the active database",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := c.ListSchemas(
				context.Background(),
				&client.ListSchemasParams{
					Instance: instancePointer(instance),
				},
			)
			errors.PanicOnError(e)
			printResponse(r)
		},
	}
	result.Flags().StringVar(
		&instance,
		"instance",
		"",
		"Instance name (optional when only one instance is configured)",
	)

	return result
}

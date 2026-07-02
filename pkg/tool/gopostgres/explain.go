package gopostgres

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/generated/client"
	"github.com/spf13/cobra"
)

func explain(c *client.Client) *cobra.Command {
	var instance string
	var analyze bool
	result := &cobra.Command{
		Use:   "explain [sql]",
		Short: "Show execution plan for a SQL query",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			r, e := c.Explain(
				context.Background(),
				client.ExplainJSONRequestBody{
					Instance: instancePointer(instance),
					Sql:      arguments[0],
					Analyze:  &analyze,
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
	result.Flags().BoolVar(
		&analyze,
		"analyze",
		false,
		"Run EXPLAIN ANALYZE",
	)

	return result
}

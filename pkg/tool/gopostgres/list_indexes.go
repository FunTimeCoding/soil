package gopostgres

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"github.com/spf13/cobra"
)

func listIndexes(c *client.Client) *cobra.Command {
	var instance string
	var schema string
	var table string
	result := &cobra.Command{
		Use:   "list-indexes",
		Short: "List indexes on a table",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			p := &client.ListIndexesParams{
				Instance: instancePointer(instance),
				Table:    table,
			}

			if schema != "" {
				p.Schema = &schema
			}

			r, e := c.ListIndexes(context.Background(), p)
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
	result.Flags().StringVar(
		&table,
		"table",
		"",
		"Table name (required)",
	)
	result.Flags().StringVar(
		&schema,
		"schema",
		"",
		"Schema name (default: public)",
	)
	errors.PanicOnError(result.MarkFlagRequired("table"))

	return result
}

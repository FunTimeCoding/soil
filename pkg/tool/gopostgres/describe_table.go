package gopostgres

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"github.com/spf13/cobra"
)

func describeTable(c *client.Client) *cobra.Command {
	var instance string
	var schema string
	result := &cobra.Command{
		Use:   "describe-table [table]",
		Short: "Show columns, types, and constraints for a table",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			p := &client.DescribeTableParams{
				Instance: instancePointer(instance),
			}

			if schema != "" {
				p.Schema = &schema
			}

			r, e := c.DescribeTable(
				context.Background(),
				arguments[0],
				p,
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
	result.Flags().StringVar(
		&schema,
		"schema",
		"",
		"Schema name (default: public)",
	)

	return result
}

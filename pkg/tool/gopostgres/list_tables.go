package gopostgres

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/face"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"github.com/spf13/cobra"
)

func listTables(
	c *face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var instance string
	var schema string
	result := &cobra.Command{
		Use:   "list-tables",
		Short: "List tables in a schema",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			p := &client.ListTablesParams{Instance: instancePointer(instance)}

			if schema != "" {
				p.Schema = &schema
			}

			r, e := (*c).ListTables(context.Background(), p)
			t.EmitValue(r, e)
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

package gopostgres

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/face"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"github.com/spf13/cobra"
)

func query(
	c *face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var instance string
	result := &cobra.Command{
		Use:   "query [sql]",
		Short: "Execute SQL against a PostgreSQL instance",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			r, e := (*c).Query(
				context.Background(),
				client.QueryJSONRequestBody{
					Instance: instancePointer(instance),
					Sql:      arguments[0],
				},
			)
			t.EmitValue(r, e)
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

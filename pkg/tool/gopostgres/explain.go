package gopostgres

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/face"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"github.com/spf13/cobra"
)

func explain(
	c *face.Client,
	t *terminal.Terminal,
) *cobra.Command {
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
			r, e := (*c).Explain(
				context.Background(),
				client.ExplainJSONRequestBody{
					Instance: instancePointer(instance),
					Sql:      arguments[0],
					Analyze:  &analyze,
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
	result.Flags().BoolVar(&analyze, "analyze", false, "Run EXPLAIN ANALYZE")

	return result
}

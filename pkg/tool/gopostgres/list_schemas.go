package gopostgres

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/face"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"github.com/spf13/cobra"
)

func listSchemas(
	c *face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var instance string
	result := &cobra.Command{
		Use:   "list-schemas",
		Short: "List schemas in the active database",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := (*c).ListSchemas(
				context.Background(),
				&client.ListSchemasParams{Instance: instancePointer(instance)},
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

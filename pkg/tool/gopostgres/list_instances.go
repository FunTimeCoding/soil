package gopostgres

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/face"
	"github.com/spf13/cobra"
)

func listInstances(
	c *face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "list-instances",
		Short: "List configured PostgreSQL instances",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := (*c).ListInstances(context.Background())
			t.EmitValue(r, e)
		},
	}
}

package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func listManufacturers(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "list-manufacturers",
		Short: "List all NetBox manufacturers",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.EmitValue(c.ListManufacturers(context.Background()))
		},
	}
}

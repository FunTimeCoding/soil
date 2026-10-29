package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func listCables(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "list-cables",
		Short: "List all NetBox cables",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.EmitValue(c.ListCables(context.Background()))
		},
	}
}

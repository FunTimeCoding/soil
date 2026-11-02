package godirectory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/face"
	"github.com/spf13/cobra"
)

func groupDelete(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "delete <name>",
		Short: "Remove a directory group",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			t.EmitValue(c.DeleteGroupName(context.Background(), a[0]))
		},
	}
}

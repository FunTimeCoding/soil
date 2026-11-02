package godirectory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/face"
	"github.com/spf13/cobra"
)

func groupRemove(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "remove <name> <account>",
		Short: "Remove a user from a group",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			t.EmitValue(
				c.DeleteGroupNameMemberAccount(context.Background(), a[0], a[1]),
			)
		},
	}
}

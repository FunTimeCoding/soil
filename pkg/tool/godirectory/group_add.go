package godirectory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/face"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/spf13/cobra"
)

func groupAdd(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "add <name> <account>",
		Short: "Add a user to a group",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			t.EmitValue(
				c.PostGroupNameMember(
					context.Background(),
					a[0],
					client.PostGroupNameMemberJSONRequestBody{Account: a[1]},
				),
			)
		},
	}
}

package godirectory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/face"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/spf13/cobra"
)

func userPassword(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "password <account> <password>",
		Short: "Set a directory user's password",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			t.EmitValue(
				c.PostUserAccountPassword(
					context.Background(),
					a[0],
					client.PostUserAccountPasswordJSONRequestBody{
						Password: a[1],
					},
				),
			)
		},
	}
}

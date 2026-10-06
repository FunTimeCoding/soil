package godirectory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/spf13/cobra"
)

func userDelete(
	c *client.ClientWithResponses,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "delete <account>",
		Short: "Remove a directory user",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			result, e := c.DeleteUserAccountWithResponse(
				context.Background(),
				a[0],
			)
			errors.PanicOnError(e)

			if result.JSON200 == nil {
				t.Reject(result.Status(), result.Body)
			}

			console.Line(notation.MarshalIndent(result.JSON200))
		},
	}
}

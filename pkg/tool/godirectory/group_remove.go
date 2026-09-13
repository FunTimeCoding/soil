package godirectory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/spf13/cobra"
)

func groupRemove(c *client.ClientWithResponses) *cobra.Command {
	return &cobra.Command{
		Use:   "remove <name> <account>",
		Short: "Remove a user from a group",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			result, e := c.DeleteGroupNameMemberAccountWithResponse(
				context.Background(),
				a[0],
				a[1],
			)
			errors.PanicOnError(e)
			console.Line(notation.MarshalIndent(result.JSON200))
		},
	}
}

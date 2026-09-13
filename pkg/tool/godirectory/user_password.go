package godirectory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/spf13/cobra"
)

func userPassword(c *client.ClientWithResponses) *cobra.Command {
	return &cobra.Command{
		Use:   "password <account> <password>",
		Short: "Set a directory user's password",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			result, e := c.PostUserAccountPasswordWithResponse(
				context.Background(),
				a[0],
				client.PostUserAccountPasswordJSONRequestBody{Password: a[1]},
			)
			errors.PanicOnError(e)
			console.Line(notation.MarshalIndent(result.JSON200))
		},
	}
}

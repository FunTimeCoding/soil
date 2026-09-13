package godirectory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/spf13/cobra"
)

func groupCreate(c *client.ClientWithResponses) *cobra.Command {
	return &cobra.Command{
		Use:   "create <name>",
		Short: "Create a directory group",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			result, e := c.PostGroupWithResponse(
				context.Background(),
				client.PostGroupJSONRequestBody{Name: a[0]},
			)
			errors.PanicOnError(e)
			console.Line(notation.MarshalIndent(result.JSON200))
		},
	}
}

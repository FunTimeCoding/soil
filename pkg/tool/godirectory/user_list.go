package godirectory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/spf13/cobra"
)

func userList(c *client.ClientWithResponses) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List directory users",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			result, e := c.GetUserWithResponse(context.Background())
			errors.PanicOnError(e)
			console.Line(notation.MarshalIndent(result.JSON200))
		},
	}
}

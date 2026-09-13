package godirectory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/spf13/cobra"
)

func groupList(c *client.ClientWithResponses) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List directory groups",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			result, e := c.GetGroupWithResponse(context.Background())
			errors.PanicOnError(e)
			console.Line(notation.MarshalIndent(result.JSON200))
		},
	}
}

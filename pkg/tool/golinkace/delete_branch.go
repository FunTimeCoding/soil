package golinkace

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/golinkace/command_context"
	"github.com/spf13/cobra"
	"strconv"
)

func deleteBranch(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-branch <list-id>",
		Short: "Delete a list and all its links",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier, e := strconv.Atoi(arguments[0])
			errors.PanicOnError(e)
			c.Terminal().EmitDone(
				c.Client().DeleteBranch(context.Background(), int32(identifier)),
			)
		},
	}
}

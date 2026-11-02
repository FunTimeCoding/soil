package golinkace

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/golinkace/command_context"
	"github.com/spf13/cobra"
	"strconv"
)

func deleteTag(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-tag <tag-id>",
		Short: "Delete a tag",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier, e := strconv.Atoi(arguments[0])
			errors.PanicOnError(e)
			c.Terminal().EmitDone(
				c.Client().DeleteTag(context.Background(), int32(identifier)),
			)
		},
	}
}

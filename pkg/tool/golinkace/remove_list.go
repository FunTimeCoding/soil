package golinkace

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/golinkace/command_context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/spf13/cobra"
	"strconv"
)

func removeList(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "remove-list <link-id> <list-name>",
		Short: "Remove a list from a link",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier, e := strconv.Atoi(arguments[0])
			errors.PanicOnError(e)
			c.Terminal().EmitValue(
				c.Client().RemoveList(
					context.Background(),
					int32(identifier),
					client.RemoveListJSONRequestBody{Name: arguments[1]},
				),
			)
		},
	}
}

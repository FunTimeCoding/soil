package golinkace

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/golinkace/command_context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/spf13/cobra"
	"strconv"
)

func appendTag(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "append-tag <link-id> <tag-name>",
		Short: "Add a tag to a link",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier, e := strconv.Atoi(arguments[0])
			errors.PanicOnError(e)
			c.Terminal().EmitValue(
				c.Client().AppendTag(
					context.Background(),
					int32(identifier),
					client.AppendTagJSONRequestBody{Name: arguments[1]},
				),
			)
		},
	}
}

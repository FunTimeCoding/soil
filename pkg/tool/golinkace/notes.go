package golinkace

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/golinkace/command_context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/spf13/cobra"
	"strconv"
)

func notes(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "notes <link-id>",
		Short: "List notes for a link",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier, e := strconv.Atoi(arguments[0])
			errors.PanicOnError(e)
			c.Terminal().EmitValue(
				c.Client().GetLinkNotes(
					context.Background(),
					int32(identifier),
					&client.GetLinkNotesParams{},
				),
			)
		},
	}
}

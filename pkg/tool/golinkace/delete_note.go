package golinkace

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/golinkace/command_context"
	"github.com/spf13/cobra"
	"strconv"
)

func deleteNote(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-note <note-id>",
		Short: "Delete a note",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier, e := strconv.Atoi(arguments[0])
			errors.PanicOnError(e)
			c.Terminal().EmitDone(
				c.Client().DeleteNote(context.Background(), int32(identifier)),
			)
		},
	}
}

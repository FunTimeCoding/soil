package golinkace

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/golinkace/command_context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/spf13/cobra"
	"strconv"
)

func editNote(c *command_context.Context) *cobra.Command {
	var message string
	result := &cobra.Command{
		Use:   "edit-note <note-id>",
		Short: "Edit a note",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier, e := strconv.Atoi(arguments[0])
			errors.PanicOnError(e)
			c.Terminal().EmitValue(
				c.Client().EditNote(
					context.Background(),
					int32(identifier),
					client.EditNoteJSONRequestBody{Text: message},
				),
			)
		},
	}
	result.Flags().StringVar(&message, "message", "", "note text")
	errors.PanicOnError(result.MarkFlagRequired("message"))

	return result
}

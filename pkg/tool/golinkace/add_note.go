package golinkace

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/golinkace/command_context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/spf13/cobra"
	"strconv"
)

func addNote(c *command_context.Context) *cobra.Command {
	var message string
	result := &cobra.Command{
		Use:   "add-note <link-id>",
		Short: "Add a note to a link",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier, e := strconv.Atoi(arguments[0])
			errors.PanicOnError(e)
			c.Terminal().EmitValue(
				c.Client().AddNote(
					context.Background(),
					client.AddNoteJSONRequestBody{
						LinkIdentifier: int32(identifier),
						Text:           message,
					},
				),
			)
		},
	}
	result.Flags().StringVar(&message, "message", "", "note text")
	errors.PanicOnError(result.MarkFlagRequired("message"))

	return result
}

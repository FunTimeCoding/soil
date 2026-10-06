package golinkace

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/linkace/note"
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
			r, f := c.Client().AddNoteWithResponse(
				context.Background(),
				client.AddNoteJSONRequestBody{
					LinkIdentifier: int32(identifier),
					Text:           message,
				},
			)
			errors.PanicOnError(f)

			if r.JSON200 == nil {
				c.Terminal().Reject(r.Status(), r.Body)
			}

			fmt.Println(note.FromDaemon(*r.JSON200).Format())
		},
	}
	result.Flags().StringVar(&message, "message", "", "note text")
	errors.PanicOnError(result.MarkFlagRequired("message"))

	return result
}

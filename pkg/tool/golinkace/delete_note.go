package golinkace

import (
	"context"
	"fmt"
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
			_, f := c.Client().DeleteNoteWithResponse(
				context.Background(),
				int32(identifier),
			)
			errors.PanicOnError(f)
			fmt.Printf("deleted note %d\n", identifier)
		},
	}
}

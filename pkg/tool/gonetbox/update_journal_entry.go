package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
	"strconv"
)

func updateJournalEntry(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var kind string
	var comments string
	result := &cobra.Command{
		Use:   "update-journal-entry [identifier]",
		Short: "Update a journal entry's comments or kind",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier, e := strconv.Atoi(arguments[0])
			errors.PanicOnError(e)
			t.EmitValue(
				c.UpdateJournalEntry(
					context.Background(),
					int32(identifier),
					kind,
					comments,
				),
			)
		},
	}
	result.Flags().StringVar(
		&comments,
		"comments",
		"",
		"new entry text (optional)",
	)
	result.Flags().StringVar(
		&kind,
		"kind",
		"",
		"new entry kind: info, success, warning, or danger (optional)",
	)

	return result
}

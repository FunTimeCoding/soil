package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func addVirtualJournalEntry(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var kind string
	result := &cobra.Command{
		Use:   "add-virtual-journal-entry [machine] [comments]",
		Short: "Add a journal entry to a virtual machine",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(
				c.AddVirtualJournalEntry(
					context.Background(),
					arguments[0],
					kind,
					arguments[1],
				),
			)
		},
	}
	result.Flags().StringVar(
		&kind,
		"kind",
		"",
		"entry kind: info, success, warning, or danger (optional)",
	)

	return result
}

package goatlassian

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/face"
	"github.com/spf13/cobra"
)

func addChecklistItem(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "add-checklist-item [key] [text]",
		Short: "Append an item to an issue's checklist",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(
				c.AddChecklistItem(
					context.Background(),
					arguments[0],
					arguments[1],
				),
			)
		},
	}
}

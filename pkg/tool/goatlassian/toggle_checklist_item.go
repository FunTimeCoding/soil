package goatlassian

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/face"
	"github.com/spf13/cobra"
	"strconv"
)

func toggleChecklistItem(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "toggle-checklist-item [key] [index]",
		Short: "Toggle a checklist item by one-based index",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			index, e := strconv.Atoi(arguments[1])

			if e != nil {
				t.Exitf("invalid index: %s\n", arguments[1])
			}

			t.EmitValue(
				c.ToggleChecklistItem(context.Background(), arguments[0], index),
			)
		},
	}
}

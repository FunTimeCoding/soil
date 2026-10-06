package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
	"strconv"
)

func toggleChecklistItem(
	c *client.Client,
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

			t.Emit(c.ToggleChecklistItem(arguments[0], index))
		},
	}
}

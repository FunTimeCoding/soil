package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
	"strconv"
)

func editChecklistItem(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "edit-checklist-item [key] [index] [text]",
		Short: "Edit a checklist item's text by one-based index",
		Args:  cobra.ExactArgs(3),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			index, e := strconv.Atoi(arguments[1])

			if e != nil {
				t.Exitf("invalid index: %s\n", arguments[1])
			}

			t.Emit(c.EditChecklistItem(arguments[0], index, arguments[2]))
		},
	}
}

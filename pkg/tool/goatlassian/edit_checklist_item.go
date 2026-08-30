package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func editChecklistItem(c *client.Client) *cobra.Command {
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
				errors.Printf("invalid index: %s\n", arguments[1])
				os.Exit(1)
			}

			console.Emit(c.EditChecklistItem(arguments[0], index, arguments[2]))
		},
	}
}

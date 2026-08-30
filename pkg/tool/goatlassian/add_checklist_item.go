package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func addChecklistItem(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "add-checklist-item [key] [text]",
		Short: "Append an item to an issue's checklist",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(c.AddChecklistItem(arguments[0], arguments[1]))
		},
	}
}

package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func deleteLink(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-link [identifier]",
		Short: "Delete an issue link",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(c.DeleteLink(arguments[0]))
			console.Format("deleted link %s\n", arguments[0])
		},
	}
}

package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func deleteComment(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-comment [key] [comment-id]",
		Short: "Delete a Jira issue comment",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.Emit(c.DeleteComment(arguments[0], arguments[1]))
			console.Format("deleted comment %s\n", arguments[1])
		},
	}
}

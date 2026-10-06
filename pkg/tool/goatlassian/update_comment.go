package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func updateComment(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "update-comment [key] [comment-id] [body]",
		Short: "Update a Jira issue comment",
		Args:  cobra.ExactArgs(3),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.Emit(c.UpdateComment(arguments[0], arguments[1], arguments[2]))
			console.Format("updated comment %s\n", arguments[1])
		},
	}
}

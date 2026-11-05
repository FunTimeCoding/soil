package goatlassian

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/face"
	"github.com/spf13/cobra"
)

func deleteComment(
	c face.Client,
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
			t.EmitDone(
				c.DeleteComment(
					context.Background(),
					arguments[0],
					arguments[1],
				),
			)
		},
	}
}

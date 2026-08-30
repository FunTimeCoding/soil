package goatlassian

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func deleteComment(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-comment [key] [comment-id]",
		Short: "Delete a Jira issue comment",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(c.DeleteComment(arguments[0], arguments[1]))
			fmt.Printf("deleted comment %s\n", arguments[1])
		},
	}
}

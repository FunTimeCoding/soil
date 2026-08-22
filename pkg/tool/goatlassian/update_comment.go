package goatlassian

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func updateComment(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "update-comment [key] [comment-id] [body]",
		Short: "Update a Jira issue comment",
		Args:  cobra.ExactArgs(3),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			c.UpdateComment(arguments[0], arguments[1], arguments[2])
			fmt.Printf("updated comment %s\n", arguments[1])
		},
	}
}

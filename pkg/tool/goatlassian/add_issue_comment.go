package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func addIssueComment(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "add-issue-comment [key] [body]",
		Short: "Add a comment to a Jira issue",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(c.AddIssueComment(arguments[0], arguments[1]))
		},
	}
}

package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func getIssue(c *client.Client) *cobra.Command {
	var comments bool
	result := &cobra.Command{
		Use:   "get-issue [key]",
		Short: "Get a Jira issue by key",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(c.GetIssue(arguments[0], comments))
		},
	}
	result.Flags().BoolVar(
		&comments,
		"comments",
		false,
		"Include the issue's comments",
	)

	return result
}

package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func transitionIssue(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "transition-issue [key] [transition-identifier]",
		Short: "Transition a Jira issue to a new status",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(c.TransitionIssue(arguments[0], arguments[1]))
		},
	}
}

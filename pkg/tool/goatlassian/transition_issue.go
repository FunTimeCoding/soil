package goatlassian

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/face"
	"github.com/spf13/cobra"
)

func transitionIssue(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "transition-issue [key] [transition-identifier]",
		Short: "Transition a Jira issue to a new status",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitDone(
				c.TransitionIssue(
					context.Background(),
					arguments[0],
					arguments[1],
				),
			)
		},
	}
}

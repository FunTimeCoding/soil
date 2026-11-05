package goatlassian

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/face"
	"github.com/spf13/cobra"
)

func linkIssues(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var linkType string
	result := &cobra.Command{
		Use:   "link-issues [key] [target-key]",
		Short: "Link two Jira issues",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitDone(
				c.LinkIssues(
					context.Background(),
					arguments[0],
					arguments[1],
					linkType,
				),
			)
		},
	}
	result.Flags().StringVar(
		&linkType,
		"type",
		"",
		"link type name (default Relates)",
	)

	return result
}

package goatlassian

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func addPageComment(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "add-page-comment [identifier] [body]",
		Short: "Add a comment to a Confluence page",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitDone(
				c.AddPageComment(
					context.Background(),
					arguments[0],
					arguments[1],
				),
			)
		},
	}
}

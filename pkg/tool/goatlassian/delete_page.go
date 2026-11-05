package goatlassian

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/face"
	"github.com/spf13/cobra"
)

func deletePage(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var draft bool
	result := &cobra.Command{
		Use:   "delete-page [identifier]",
		Short: "Delete a Confluence page (or its draft)",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitDone(c.DeletePage(context.Background(), arguments[0], draft))
		},
	}
	result.Flags().BoolVar(
		&draft,
		"draft",
		false,
		"delete the draft instead of the published page",
	)

	return result
}

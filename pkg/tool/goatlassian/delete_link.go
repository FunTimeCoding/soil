package goatlassian

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/face"
	"github.com/spf13/cobra"
)

func deleteLink(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-link [identifier]",
		Short: "Delete an issue link",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitDone(c.DeleteLink(context.Background(), arguments[0]))
		},
	}
}

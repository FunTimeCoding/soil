package goatlassian

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/face"
	"github.com/spf13/cobra"
)

func listPages(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var status string
	result := &cobra.Command{
		Use:   "list-pages [space-identifier]",
		Short: "List Confluence pages in a space",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(c.ListPages(context.Background(), arguments[0], status))
		},
	}
	result.Flags().StringVar(&status, "status", "", "page status filter")

	return result
}

package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func listSpaces(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "list-spaces",
		Short: "List all visible Confluence spaces",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.Emit(c.ListSpaces())
		},
	}
}

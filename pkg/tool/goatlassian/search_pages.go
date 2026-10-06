package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func searchPages(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "search-pages [query]",
		Short: "Search Confluence pages using CQL or plain text",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.Emit(c.SearchPages(arguments[0]))
		},
	}
}

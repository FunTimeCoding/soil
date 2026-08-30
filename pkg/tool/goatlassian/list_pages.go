package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func listPages(c *client.Client) *cobra.Command {
	var status string
	result := &cobra.Command{
		Use:   "list-pages [space-identifier]",
		Short: "List Confluence pages in a space",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(c.ListPages(arguments[0], status))
		},
	}
	result.Flags().StringVar(&status, "status", "", "page status filter")

	return result
}

package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func getPageDraft(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "get-page-draft [identifier]",
		Short: "Get a page with unpublished draft changes overlaid",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(c.GetPageDraft(arguments[0]))
		},
	}
}

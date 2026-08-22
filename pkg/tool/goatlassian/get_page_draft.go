package goatlassian

import (
	"fmt"
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
			fmt.Println(c.GetPageDraft(arguments[0]))
		},
	}
}

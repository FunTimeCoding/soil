package goatlassian

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func getPageChildren(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "get-page-children [identifier]",
		Short: "List child pages of a Confluence page",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.GetPageChildren(arguments[0]))
		},
	}
}

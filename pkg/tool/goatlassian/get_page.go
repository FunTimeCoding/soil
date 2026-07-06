package goatlassian

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func getPage(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "get-page [identifier]",
		Short: "Get a Confluence page by ID with body content",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.GetPage(arguments[0]))
		},
	}
}

package goatlassian

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func addPageComment(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "add-page-comment [identifier] [body]",
		Short: "Add a comment to a Confluence page",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.AddPageComment(arguments[0], arguments[1]))
		},
	}
}

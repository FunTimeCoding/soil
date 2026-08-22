package goatlassian

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func deletePage(c *client.Client) *cobra.Command {
	var draft bool
	result := &cobra.Command{
		Use:   "delete-page [identifier]",
		Short: "Delete a Confluence page (or its draft)",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			c.DeletePage(arguments[0], draft)
			fmt.Printf("deleted page %s\n", arguments[0])
		},
	}
	result.Flags().BoolVar(
		&draft,
		"draft",
		false,
		"delete the draft instead of the published page",
	)

	return result
}

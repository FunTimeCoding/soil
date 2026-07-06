package goatlassian

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func createPage(c *client.Client) *cobra.Command {
	var space string
	var parent string
	result := &cobra.Command{
		Use:   "create-page [title] [body]",
		Short: "Create a Confluence page with markdown content",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(
				c.CreatePage(
					space,
					parent,
					arguments[0],
					arguments[1],
				),
			)
		},
	}
	result.Flags().StringVar(
		&space,
		"space",
		"",
		"space identifier (required)",
	)
	result.Flags().StringVar(
		&parent,
		"parent",
		"",
		"parent page identifier (required)",
	)
	errors.PanicOnError(result.MarkFlagRequired("space"))
	errors.PanicOnError(result.MarkFlagRequired("parent"))

	return result
}

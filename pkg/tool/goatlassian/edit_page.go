package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func editPage(c *client.Client) *cobra.Command {
	var old string
	var replacement string
	var title string
	var message string
	var draft bool
	result := &cobra.Command{
		Use:   "edit-page [identifier]",
		Short: "Replace text in a Confluence page",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(
				c.EditPage(
					arguments[0],
					old,
					replacement,
					title,
					message,
					draft,
				),
			)
		},
	}
	result.Flags().StringVar(&old, "old", "", "text to replace")
	result.Flags().StringVar(&replacement, "new", "", "replacement text")
	result.Flags().StringVar(&title, "title", "", "new page title")
	result.Flags().StringVar(&message, "message", "", "version message")
	result.Flags().BoolVar(&draft, "draft", false, "edit the draft")
	errors.PanicOnError(result.MarkFlagRequired("old"))
	errors.PanicOnError(result.MarkFlagRequired("new"))

	return result
}

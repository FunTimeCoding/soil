package golinkace

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/golinkace/command_context"
	"github.com/funtimecoding/soil/pkg/tool/golinkace/constant"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/spf13/cobra"
)

func createLink(c *command_context.Context) *cobra.Command {
	var listName string
	var title string
	result := &cobra.Command{
		Use:   "create-link <url>",
		Short: "Create a new link",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			c.Terminal().EmitValue(
				c.Client().AddLink(
					context.Background(),
					client.AddLinkJSONRequestBody{
						Link: arguments[0],
						Name: &title,
						List: &listName,
					},
				),
			)
		},
	}
	result.Flags().StringVar(
		&listName,
		constant.ListFlag,
		"",
		"list name or ID",
	)
	result.Flags().StringVar(&title, "title", "", "link title")
	errors.PanicOnError(result.MarkFlagRequired("title"))

	return result
}

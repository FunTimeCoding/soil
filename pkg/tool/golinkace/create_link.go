package golinkace

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	linkaceConstant "github.com/funtimecoding/soil/pkg/linkace/constant"
	"github.com/funtimecoding/soil/pkg/linkace/link"
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
			body := client.AddLinkJSONRequestBody{
				Link: arguments[0],
				Name: &title,
				List: &listName,
			}
			r, e := c.Client().AddLinkWithResponse(context.Background(), body)
			errors.PanicOnError(e)

			if r.JSON200 == nil {
				c.Terminal().Reject(r.Status(), r.Body)
			}

			fmt.Println(
				link.FromDaemon(*r.JSON200, c.Host()).Format(
					linkaceConstant.Format,
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

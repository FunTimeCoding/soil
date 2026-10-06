package golinkace

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/linkace/constant"
	"github.com/funtimecoding/soil/pkg/linkace/link"
	"github.com/funtimecoding/soil/pkg/tool/golinkace/command_context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/spf13/cobra"
	"strconv"
)

func edit(c *command_context.Context) *cobra.Command {
	var name string
	var l string
	var description string
	var tagList string
	var listList string
	result := &cobra.Command{
		Use:   "edit <link-id>",
		Short: "Edit a link",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier, e := strconv.Atoi(arguments[0])
			errors.PanicOnError(e)
			body := client.EditLinkJSONRequestBody{}

			if name != "" {
				body.Name = &name
			}

			if l != "" {
				body.Link = &l
			}

			if description != "" {
				body.Description = &description
			}

			if tagList != "" {
				tags := splitTrimmed(tagList)
				body.Tags = &tags
			}

			if listList != "" {
				lists := splitTrimmed(listList)
				body.Lists = &lists
			}

			r, f := c.Client().EditLinkWithResponse(
				context.Background(),
				int32(identifier),
				body,
			)
			errors.PanicOnError(f)

			if r.JSON200 == nil {
				c.Terminal().Reject(r.Status(), r.Body)
			}

			fmt.Println(
				link.FromDaemon(*r.JSON200, c.Host()).Format(constant.Format),
			)
		},
	}
	result.Flags().StringVar(&name, "name", "", "new title")
	result.Flags().StringVar(&l, "url", "", "new URL")
	result.Flags().StringVar(&description, "description", "", "new description")
	result.Flags().StringVar(
		&tagList,
		"tags",
		"",
		"comma-separated tag names (replaces all)",
	)
	result.Flags().StringVar(
		&listList,
		"lists",
		"",
		"comma-separated list names (replaces all)",
	)

	return result
}

package golinkace

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/linkace/constant"
	"github.com/funtimecoding/soil/pkg/linkace/tag"
	"github.com/funtimecoding/soil/pkg/tool/golinkace/command_context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/spf13/cobra"
	"strconv"
)

func editTag(c *command_context.Context) *cobra.Command {
	var name string
	result := &cobra.Command{
		Use:   "edit-tag <tag-id>",
		Short: "Edit a tag",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier, e := strconv.Atoi(arguments[0])
			errors.PanicOnError(e)
			body := client.EditTagJSONRequestBody{}

			if name != "" {
				body.Name = &name
			}

			r, f := c.Client().EditTagWithResponse(
				context.Background(),
				int32(identifier),
				body,
			)
			errors.PanicOnError(f)

			if r.JSON200 == nil {
				c.Terminal().Reject(r.Status(), r.Body)
			}

			fmt.Println(
				tag.FromDaemon(*r.JSON200, c.Host()).Format(constant.Format),
			)
		},
	}
	result.Flags().StringVar(&name, "name", "", "new name")
	errors.PanicOnError(result.MarkFlagRequired("name"))

	return result
}

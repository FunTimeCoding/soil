package golinkace

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/linkace/constant"
	"github.com/funtimecoding/soil/pkg/linkace/list"
	"github.com/funtimecoding/soil/pkg/tool/golinkace/command_context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/spf13/cobra"
	"strconv"
)

func editList(c *command_context.Context) *cobra.Command {
	var name string
	var description string
	result := &cobra.Command{
		Use:   "edit-list <list-id>",
		Short: "Edit a list",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier, e := strconv.Atoi(arguments[0])
			errors.PanicOnError(e)
			body := client.EditListJSONRequestBody{}

			if name != "" {
				body.Name = &name
			}

			if description != "" {
				body.Description = &description
			}

			r, f := c.Client().EditListWithResponse(
				context.Background(),
				int32(identifier),
				body,
			)
			errors.PanicOnError(f)
			fmt.Println(
				list.FromDaemon(*r.JSON200, c.Host()).Format(constant.Format),
			)
		},
	}
	result.Flags().StringVar(&name, "name", "", "new name")
	result.Flags().StringVar(&description, "description", "", "new description")

	return result
}

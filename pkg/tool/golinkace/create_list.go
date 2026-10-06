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
)

func createList(c *command_context.Context) *cobra.Command {
	var description string
	result := &cobra.Command{
		Use:   "create-list <name>",
		Short: "Create a new list",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			r, e := c.Client().CreateListWithResponse(
				context.Background(),
				client.CreateListJSONRequestBody{
					Name:        arguments[0],
					Description: &description,
				},
			)
			errors.PanicOnError(e)

			if r.JSON200 == nil {
				c.Terminal().Reject(r.Status(), r.Body)
			}

			fmt.Println(
				list.FromDaemon(*r.JSON200, c.Host()).Format(constant.Format),
			)
		},
	}
	result.Flags().StringVar(
		&description,
		"description",
		"",
		"list description",
	)

	return result
}

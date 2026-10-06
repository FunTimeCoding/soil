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
)

func createTag(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "create-tag <name>",
		Short: "Create a new tag",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			r, e := c.Client().CreateTagWithResponse(
				context.Background(),
				client.CreateTagJSONRequestBody{Name: arguments[0]},
			)
			errors.PanicOnError(e)

			if r.JSON200 == nil {
				c.Terminal().Reject(r.Status(), r.Body)
			}

			fmt.Println(
				tag.FromDaemon(*r.JSON200, c.Host()).Format(constant.Format),
			)
		},
	}
}

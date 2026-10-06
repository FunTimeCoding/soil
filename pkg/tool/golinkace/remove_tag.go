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

func removeTag(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "remove-tag <link-id> <tag-name>",
		Short: "Remove a tag from a link",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier, e := strconv.Atoi(arguments[0])
			errors.PanicOnError(e)
			r, f := c.Client().RemoveTagWithResponse(
				context.Background(),
				int32(identifier),
				client.RemoveTagJSONRequestBody{Name: arguments[1]},
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
}

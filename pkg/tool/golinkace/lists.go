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

func lists(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "lists",
		Short: "List all lists",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := c.Client().GetListsWithResponse(
				context.Background(),
				&client.GetListsParams{},
			)
			errors.PanicOnError(e)
			f := constant.Format

			for _, l := range *r.JSON200.Lists {
				fmt.Println(list.FromDaemon(l, c.Host()).Format(f))
			}
		},
	}
}

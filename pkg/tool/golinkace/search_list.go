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

func searchList(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "search-list <query>",
		Short: "Search lists by name",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			entityType := "list"
			r, e := c.Client().SearchWithResponse(
				context.Background(),
				&client.SearchParams{Query: arguments[0], Type: &entityType},
			)
			errors.PanicOnError(e)
			f := constant.Format

			if r.JSON200.Lists != nil {
				for _, l := range *r.JSON200.Lists {
					fmt.Println(list.FromDaemon(l, c.Host()).Format(f))
				}
			}
		},
	}
}

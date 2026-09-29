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

func searchTag(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "search-tag <query>",
		Short: "Search tags by name",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			entityType := "tag"
			r, e := c.Client().SearchWithResponse(
				context.Background(),
				&client.SearchParams{Query: arguments[0], Type: &entityType},
			)
			errors.PanicOnError(e)
			f := constant.Format

			if r.JSON200.Tags != nil {
				for _, t := range *r.JSON200.Tags {
					fmt.Println(tag.FromDaemon(t, c.Host()).Format(f))
				}
			}
		},
	}
}

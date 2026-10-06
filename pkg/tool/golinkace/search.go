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
)

func search(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "search <query>",
		Short: "Search links",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			entityType := "link"
			r, e := c.Client().SearchWithResponse(
				context.Background(),
				&client.SearchParams{Query: arguments[0], Type: &entityType},
			)
			errors.PanicOnError(e)

			if r.JSON200 == nil {
				c.Terminal().Reject(r.Status(), r.Body)
			}

			if r.JSON200.Links == nil || len(*r.JSON200.Links) == 0 {
				errors.Printf("no results")

				return
			}

			f := constant.Format

			for _, l := range *r.JSON200.Links {
				fmt.Println(link.FromDaemon(l, c.Host()).Format(f))
			}
		},
	}
}

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

func links(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "links",
		Short: "List all links",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := c.Client().GetLinksWithResponse(
				context.Background(),
				&client.GetLinksParams{},
			)
			errors.PanicOnError(e)
			f := constant.Format

			for _, l := range *r.JSON200.Links {
				fmt.Println(link.FromDaemon(l, c.Host()).Format(f))
			}
		},
	}
}

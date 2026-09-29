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

func tags(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "tags",
		Short: "List all tags",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := c.Client().GetTagsWithResponse(
				context.Background(),
				&client.GetTagsParams{},
			)
			errors.PanicOnError(e)
			f := constant.Format

			for _, t := range *r.JSON200.Tags {
				fmt.Println(tag.FromDaemon(t, c.Host()).Format(f))
			}
		},
	}
}

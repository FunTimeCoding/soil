package golinkace

import (
	"context"
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
			c.Terminal().EmitValue(
				c.Client().GetTags(
					context.Background(),
					&client.GetTagsParams{},
				),
			)
		},
	}
}

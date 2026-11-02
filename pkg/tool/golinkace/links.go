package golinkace

import (
	"context"
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
			c.Terminal().EmitValue(
				c.Client().GetLinks(
					context.Background(),
					&client.GetLinksParams{},
				),
			)
		},
	}
}

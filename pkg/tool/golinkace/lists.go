package golinkace

import (
	"context"
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
			c.Terminal().EmitValue(
				c.Client().GetLists(
					context.Background(),
					&client.GetListsParams{},
				),
			)
		},
	}
}

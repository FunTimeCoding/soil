package golinkace

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkace/command_context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/spf13/cobra"
)

func createTag(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "create-tag <name>",
		Short: "Create a new tag",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			c.Terminal().EmitValue(
				c.Client().CreateTag(
					context.Background(),
					client.CreateTagJSONRequestBody{Name: arguments[0]},
				),
			)
		},
	}
}

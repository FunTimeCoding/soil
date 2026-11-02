package golinkace

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkace/command_context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/spf13/cobra"
)

func createList(c *command_context.Context) *cobra.Command {
	var description string
	result := &cobra.Command{
		Use:   "create-list <name>",
		Short: "Create a new list",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			c.Terminal().EmitValue(
				c.Client().CreateList(
					context.Background(),
					client.CreateListJSONRequestBody{
						Name:        arguments[0],
						Description: &description,
					},
				),
			)
		},
	}
	result.Flags().StringVar(
		&description,
		"description",
		"",
		"list description",
	)

	return result
}

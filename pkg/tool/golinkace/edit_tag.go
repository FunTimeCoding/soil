package golinkace

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/golinkace/command_context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/spf13/cobra"
	"strconv"
)

func editTag(c *command_context.Context) *cobra.Command {
	var name string
	result := &cobra.Command{
		Use:   "edit-tag <tag-id>",
		Short: "Edit a tag",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier, e := strconv.Atoi(arguments[0])
			errors.PanicOnError(e)
			body := client.EditTagJSONRequestBody{}

			if name != "" {
				body.Name = &name
			}

			c.Terminal().EmitValue(
				c.Client().EditTag(
					context.Background(),
					int32(identifier),
					body,
				),
			)
		},
	}
	result.Flags().StringVar(&name, "name", "", "new name")
	errors.PanicOnError(result.MarkFlagRequired("name"))

	return result
}

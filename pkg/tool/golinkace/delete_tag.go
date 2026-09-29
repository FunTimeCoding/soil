package golinkace

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/golinkace/command_context"
	"github.com/spf13/cobra"
	"strconv"
)

func deleteTag(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-tag <tag-id>",
		Short: "Delete a tag",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier, e := strconv.Atoi(arguments[0])
			errors.PanicOnError(e)
			_, f := c.Client().DeleteTagWithResponse(
				context.Background(),
				int32(identifier),
			)
			errors.PanicOnError(f)
			fmt.Printf("deleted tag %d\n", identifier)
		},
	}
}

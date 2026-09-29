package golinkace

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/golinkace/command_context"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func deleteLink(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "delete <link-id>",
		Short: "Delete a link",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier, e := strconv.Atoi(arguments[0])
			errors.PanicOnError(e)
			r, f := c.Client().DeleteLinkWithResponse(
				context.Background(),
				int32(identifier),
			)
			errors.PanicOnError(f)

			if r.HTTPResponse.StatusCode != 204 {
				errors.Printf(
					"unexpected status: %s\n%s\n",
					r.HTTPResponse.Status,
					string(r.Body),
				)
				os.Exit(1)
			}

			fmt.Printf("deleted link %d\n", identifier)
		},
	}
}

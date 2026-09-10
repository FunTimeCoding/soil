package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/spf13/cobra"
)

func sessionDelete(c *command_context.Context) *cobra.Command {
	var confirm string
	result := &cobra.Command{
		Use:   "delete <identifier>",
		Short: "Delete a session and its data",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			parameters := &client.DeleteSessionByIdParams{}

			if confirm != "" {
				parameters.Confirm = &confirm
			}

			response, e := c.Client().DeleteSessionByIdWithResponse(
				context.Background(),
				arguments[0],
				parameters,
			)
			errors.PanicOnError(e)

			if response.JSON409 != nil {
				console.Format("refused: %s\n", response.JSON409.Error)

				return
			}

			if response.JSON200 == nil {
				console.Format(
					"unexpected response: %s\n",
					response.HTTPResponse.Status,
				)

				return
			}

			printDeleteReceipt(response.JSON200)
		},
	}
	result.Flags().StringVar(
		&confirm,
		"confirm",
		"",
		"Confirmation hash from the session detail page",
	)

	return result
}

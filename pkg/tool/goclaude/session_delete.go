package goclaude

import (
	"context"
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
			parameters := &client.DeleteSessionByIdentifierParams{}

			if confirm != "" {
				parameters.Confirm = &confirm
			}

			r, e := c.Client().DeleteSessionByIdentifier(
				context.Background(),
				arguments[0],
				parameters,
			)

			if e != nil {
				c.Terminal().Exitln(e)

				return
			}

			printDeleteReceipt(r)
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

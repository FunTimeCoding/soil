package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/spf13/cobra"
)

func sessionDelete(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "delete <id-or-name>",
		Short: "Delete a session and its data",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier := resolveSession(c.Client(), arguments[0])

			if identifier == "" {
				console.Format("session not found: %s\n", arguments[0])

				return
			}

			_, e := c.Client().DeleteSessionByIdWithResponse(
				context.Background(),
				identifier,
			)
			errors.PanicOnError(e)
			console.Format("deleted: %s\n", identifier)
		},
	}
}

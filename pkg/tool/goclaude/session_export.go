package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/spf13/cobra"
)

func sessionExport(c *command_context.Context) *cobra.Command {
	var all bool
	result := &cobra.Command{
		Use:   "export [id-or-name]",
		Short: "Export a session to markdown",
		Args:  cobra.MaximumNArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			if all {
				c.Terminal().EmitValue(
					c.Client().PostSessionsExport(context.Background()),
				)

				return
			}

			if len(arguments) == 0 {
				c.Terminal().EmitDone(
					validation.New("provide a session id or use --all"),
				)

				return
			}

			identifier := resolveSession(c, arguments[0])

			if identifier == "" {
				return
			}

			c.Terminal().EmitValue(
				c.Client().PostSessionExport(context.Background(), identifier),
			)
		},
	}
	result.Flags().BoolVar(&all, "all", false, "Export all sessions")

	return result
}

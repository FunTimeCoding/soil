package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
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
				r, e := c.Client().PostSessionsExport(context.Background())

				if e != nil {
					c.Terminal().Exitln(e)

					return
				}

				for _, path := range r.Paths {
					console.Format("exported: %s\n", path)
				}

				return
			}

			if len(arguments) == 0 {
				console.Line("provide a session id or use --all")

				return
			}

			identifier := resolveSession(c, arguments[0])

			if identifier == "" {
				console.Format("session not found: %s\n", arguments[0])

				return
			}

			r, e := c.Client().PostSessionExport(
				context.Background(),
				identifier,
			)

			if e != nil {
				c.Terminal().Exitln(e)

				return
			}

			for _, path := range r.Paths {
				console.Format("exported: %s\n", path)
			}
		},
	}
	result.Flags().BoolVar(&all, "all", false, "Export all sessions")

	return result
}

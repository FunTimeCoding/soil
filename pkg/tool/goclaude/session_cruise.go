package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/spf13/cobra"
)

func sessionCruise(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "cruise <id-or-name> [off|default|<minutes>]",
		Short: "Show a session's cruise grace, or set it",
		Args:  cobra.RangeArgs(1, 2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier := resolveSession(c, arguments[0])

			if identifier == "" {
				return
			}

			if len(arguments) == 1 {
				c.Terminal().EmitValue(
					c.Client().GetSessionCruise(
						context.Background(),
						identifier,
					),
				)

				return
			}

			c.Terminal().EmitValue(
				c.Client().PostSessionCruise(
					context.Background(),
					identifier,
					client.CruiseRequest{Grace: arguments[1]},
				),
			)
		},
	}
}

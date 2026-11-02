package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/spf13/cobra"
)

func sessionCruise(c *command_context.Context) *cobra.Command {
	var pace int
	result := &cobra.Command{
		Use:   "cruise <id-or-name> [off|on|paced]",
		Short: "Show a session's cruise setting, or set it",
		Args:  cobra.RangeArgs(1, 2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier := resolveSession(c, arguments[0])

			if identifier == "" {
				console.Format("session not found: %s\n", arguments[0])

				return
			}

			var r *client.CruiseResponse
			var e error

			if len(arguments) == 1 {
				r, e = c.Client().GetSessionCruise(
					context.Background(),
					identifier,
				)
			} else {
				body := client.CruiseRequest{
					Mode: client.CruiseRequestMode(arguments[1]),
				}

				if pace > 0 {
					body.PaceMinutes = &pace
				}

				r, e = c.Client().PostSessionCruise(
					context.Background(),
					identifier,
					body,
				)
			}

			if e != nil {
				c.Terminal().Exitln(e)

				return
			}

			console.Line(CruiseSummary(r))
		},
	}
	result.Flags().IntVar(
		&pace,
		constant.Pace,
		0,
		"minutes between wakes, for paced",
	)

	return result
}

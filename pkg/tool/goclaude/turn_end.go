package goclaude

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/spf13/cobra"
)

func turnEnd(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "turn-end",
		Short: "Record the turn boundary in goclauded (Stop hook)",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			input := readHookInput()

			if input.SessionIdentifier == "" {
				return
			}

			RunTurnEnd(c.Client(), input.SessionIdentifier)
		},
	}
}

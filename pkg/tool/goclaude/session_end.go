package goclaude

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/spf13/cobra"
)

func sessionEnd(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "session-end",
		Short: "Mark session closed in goclauded (SessionEnd hook)",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			input := readHookInput()

			if input.SessionIdentifier == "" {
				return
			}

			RunSessionEnd(c.Client(), input.SessionIdentifier, input.Reason)
		},
	}
}

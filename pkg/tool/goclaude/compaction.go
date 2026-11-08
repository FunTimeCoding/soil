package goclaude

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/spf13/cobra"
)

func compaction(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "compaction",
		Short: "Record a compaction starting or finishing in goclauded (PreCompact and PostCompact hooks)",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			input := readHookInput()
			phase, known := CompactionPhase(input.HookEventName)

			if input.SessionIdentifier == "" || !known {
				return
			}

			RunCompaction(
				c.Client(),
				input.SessionIdentifier,
				phase,
				input.Trigger,
			)
		},
	}
}

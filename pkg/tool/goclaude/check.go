package goclaude

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/spf13/cobra"
	"os"
)

func check(c *command_context.Context) *cobra.Command {
	result := &cobra.Command{
		Use:   "check",
		Short: "Check goclauded for session roster and messages (hook output)",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			o := RunCheck(
				c.Client(),
				sessionIdentifierFromEnvironment(),
			)

			if o == "" {
				return
			}

			errors.PanicOnError(
				json.NewEncoder(os.Stdout).Encode(
					hookOutput{
						HookSpecificOutput: hookSpecificOutput{
							HookEventName:     "UserPromptSubmit",
							AdditionalContext: o,
						},
					},
				),
			)
		},
	}

	return result
}

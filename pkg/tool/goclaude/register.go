package goclaude

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/spf13/cobra"
	"os"
)

func register(c *command_context.Context) *cobra.Command {
	result := &cobra.Command{
		Use:   "register",
		Short: "Register session with goclauded (SessionStart hook)",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			session := readHookInput().SessionID

			if session == "" {
				return
			}

			name := RunRegister(c.Client(), session)

			if name == "" {
				return
			}

			environmentFile := os.Getenv(constant.EnvironmentFileEnvironment)

			if environmentFile != "" {
				f, fe := os.OpenFile(
					environmentFile,
					os.O_APPEND|os.O_CREATE|os.O_WRONLY,
					0644,
				)

				if fe == nil {
					writer.Print(
						f,
						"export %s=%s\n",
						constant.SessionIdentifierEnvironment,
						session,
					)
					writer.Print(
						f,
						"export %s=%s\n",
						constant.NameEnvironment,
						name,
					)
					errors.PanicOnError(f.Close())
				}
			}

			errors.PanicOnError(
				json.NewEncoder(os.Stdout).Encode(
					hookOutput{
						HookSpecificOutput: hookSpecificOutput{
							HookEventName: "SessionStart",
							AdditionalContext: fmt.Sprintf(
								"[goclauded] Called %s today.",
								name,
							),
						},
					},
				),
			)
		},
	}

	return result
}

package goclaude

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
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
			input := readHookInput()
			session := input.SessionIdentifier

			if session == "" {
				return
			}

			name := RunRegister(c.Client(), session)

			if name == "" {
				return
			}

			environmentFile := os.Getenv(constant.EnvironmentFileEnvironment)

			if environmentFile != "" {
				f, e := os.OpenFile(
					environmentFile,
					os.O_APPEND|os.O_CREATE|os.O_WRONLY,
					0644,
				)

				if e == nil {
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

			lines := []string{fmt.Sprintf("[goclauded] Called %s today.", name)}
			warning := labelWarning(input.WorkingDirectory)

			if warning != "" {
				lines = append(lines, warning)
			}

			errors.PanicOnError(
				json.NewEncoder(os.Stdout).Encode(
					HookOutput{
						SystemMessage: warning,
						HookSpecificOutput: HookSpecificOutput{
							HookEventName:     "SessionStart",
							AdditionalContext: join.NewLine(lines),
						},
					},
				),
			)
		},
	}

	return result
}

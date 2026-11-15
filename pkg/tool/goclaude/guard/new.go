package guard

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/response"
	"github.com/spf13/cobra"
	"runtime"
)

func New(t *terminal.Terminal) *cobra.Command {
	result := &cobra.Command{
		Use:   "guard",
		Short: "Check a tool call for command mistakes (PreToolUse hook)",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			if environment.Exists(constant.NoGuardEnvironment) {
				return
			}

			i := readInput()

			if i.ToolName != "Bash" {
				return
			}

			if b := BlockReason(runtime.GOOS, i.ToolInput.Command); b != "" {
				t.Blockln(constant.GuardBlockExit, b)

				return
			}

			a := AskReason(i.ToolInput.Command)
			h := Hint(i.ToolInput.Command)

			if a != "" || h != "" {
				t.Line(string(notation.Marshal(response.NewDecision(a, h))))
			}
		},
	}

	return result
}

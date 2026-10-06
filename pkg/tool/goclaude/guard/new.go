package guard

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
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

			v := Verdict(runtime.GOOS, i.ToolInput.Command)

			if v == "" {
				return
			}

			t.Blockln(constant.GuardBlockExit, v)
		},
	}

	return result
}

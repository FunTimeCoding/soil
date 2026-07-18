package guard

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/spf13/cobra"
	"os"
	"runtime"
)

func New() *cobra.Command {
	result := &cobra.Command{
		Use:   "guard",
		Short: "Check a tool call for command mistakes (PreToolUse hook)",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			i := readInput()

			if i.ToolName != "Bash" {
				return
			}

			v := verdict(runtime.GOOS, i.ToolInput.Command)

			if v == "" {
				return
			}

			errors.Printf("%s\n", v)
			os.Exit(constant.GuardBlockExit)
		},
	}

	return result
}

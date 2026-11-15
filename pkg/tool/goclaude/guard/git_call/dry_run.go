package git_call

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"
	"slices"
)

func (c *Call) dryRun() bool {
	for _, a := range c.arguments {
		w, literal := shell_script.Argument(a)

		if literal && w == "--" {
			return false
		}

		if literal && slices.Contains(constant.GitDryRunOptions, w) {
			return true
		}
	}

	return false
}

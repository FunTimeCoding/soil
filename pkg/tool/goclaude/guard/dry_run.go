package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"slices"
)

func (c *GitCall) dryRun() bool {
	for _, a := range c.arguments {
		w, literal := literalWord(a)

		if literal && w == "--" {
			return false
		}

		if literal && slices.Contains(constant.GitDryRunOptions, w) {
			return true
		}
	}

	return false
}

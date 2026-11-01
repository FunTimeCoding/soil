package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"slices"
)

func (c *GitCall) stashVerb() string {
	if c.subcommand != constant.GitStashSubcommand || len(c.arguments) == 0 {
		return ""
	}

	w, literal := literalWord(c.arguments[0])

	if !literal || !slices.Contains(constant.GitStashVerbs, w) {
		return ""
	}

	return w
}

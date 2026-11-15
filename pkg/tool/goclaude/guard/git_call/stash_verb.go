package git_call

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"
	"slices"
)

func (c *Call) stashVerb() string {
	if c.subcommand != constant.GitStashSubcommand || len(c.arguments) == 0 {
		return ""
	}

	w, literal := shell_script.Argument(c.arguments[0])

	if !literal || !slices.Contains(constant.GitStashVerbs, w) {
		return ""
	}

	return w
}

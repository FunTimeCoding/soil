package git_call

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"
	"slices"
	"strings"
)

func (c *Call) Grantable() bool {
	if c.hidden || !c.readable || c.fromInput || c.redirected || !c.topLevel {
		return false
	}

	if !slices.Contains(constant.GitGrantedWriters, c.subcommand) {
		return false
	}

	push := c.subcommand == constant.GitPushSubcommand
	positional := 0

	for _, w := range c.arguments {
		value, literal := shell_script.Argument(w)

		if !literal {
			return false
		}

		if !push {
			continue
		}

		if refusedPushArgument(value) {
			return false
		}

		if strings.HasPrefix(value, "-") {
			continue
		}

		positional++

		if positional == 1 && !constant.GitRemoteName.MatchString(value) {
			return false
		}
	}

	return true
}

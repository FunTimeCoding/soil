package guard

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"slices"
)

func (c *GitCall) Block() string {
	if !c.readable {
		return fmt.Sprintf(constant.GitUnreadableMessage, "subcommand")
	}

	if slices.Contains(constant.GitHistoryWriters, c.subcommand) {
		return fmt.Sprintf(constant.GitHistoryMessage, c.subcommand)
	}

	if !slices.Contains(constant.GitTreeWriters, c.subcommand) {
		return ""
	}

	if c.fromInput {
		return fmt.Sprintf(constant.GitUnreadableMessage, c.subcommand)
	}

	for _, w := range c.arguments {
		if _, literal := literalWord(w); !literal {
			return fmt.Sprintf(constant.GitUnreadableMessage, c.subcommand)
		}
	}

	if n := len(c.arguments); n > 0 && c.arguments[n-1].Lit() == "--" {
		return fmt.Sprintf(constant.GitUnreadableMessage, c.subcommand)
	}

	return ""
}

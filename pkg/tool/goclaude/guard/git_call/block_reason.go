package git_call

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"
	"slices"
)

func (c *Call) BlockReason() string {
	if c.hidden {
		return constant.GitHiddenReason
	}

	if !c.readable {
		return fmt.Sprintf(constant.GitUnreadableReason, "subcommand")
	}

	if c.listsTags() || c.readsOnly() {
		return ""
	}

	if slices.Contains(constant.GitHistoryWriters, c.subcommand) {
		return fmt.Sprintf(constant.GitHistoryReason, c.subcommand)
	}

	if !slices.Contains(constant.GitTreeWriters, c.subcommand) {
		return ""
	}

	if c.fromInput {
		return fmt.Sprintf(constant.GitUnreadableReason, c.subcommand)
	}

	for _, w := range c.arguments {
		if _, literal := shell_script.Argument(w); !literal {
			return fmt.Sprintf(constant.GitUnreadableReason, c.subcommand)
		}
	}

	if n := len(c.arguments); n > 0 && c.arguments[n-1].Lit() == "--" {
		return fmt.Sprintf(constant.GitUnreadableReason, c.subcommand)
	}

	return ""
}

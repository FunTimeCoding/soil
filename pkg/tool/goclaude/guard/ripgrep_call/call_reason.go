package ripgrep_call

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"
)

func callReason(
	c *shell_script.Call,
	configured bool,
) string {
	if c.Name != constant.RipgrepCommand {
		return ""
	}

	if c.FromInput || configured {
		return constant.RipgrepUnreadableReason
	}

	for _, w := range c.Words[1:] {
		if w.Lit() == "--" {
			return ""
		}

		if r := wordReason(w); r != "" {
			return r
		}
	}

	return ""
}

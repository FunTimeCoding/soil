package git_call

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"
	"slices"
	"strings"
)

func (c *Call) listsTags() bool {
	if c.subcommand != constant.GitTagSubcommand || c.fromInput {
		return false
	}

	listing := false
	positional := false

	for i := 0; i < len(c.arguments); i++ {
		w, literal := shell_script.Argument(c.arguments[i])

		if !literal {
			return false
		}

		name, _, valued := strings.Cut(w, "=")

		switch {
		case slices.Contains(constant.GitTagListOptions, w):
			listing = true
		case slices.Contains(constant.GitTagValueOptions, name):
			if !valued &&
				i+1 < len(c.arguments) &&
				!strings.HasPrefix(c.arguments[i+1].Lit(), "-") {
				i++
			}
		case slices.Contains(constant.GitTagReadOptions, w),
			constant.GitTagLineCount.MatchString(w):

			continue
		case strings.HasPrefix(w, "-"):
			return false
		default:
			positional = true
		}
	}

	return listing || !positional
}

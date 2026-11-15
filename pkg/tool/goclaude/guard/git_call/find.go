package git_call

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"
	"slices"
	"strings"
)

func Find(c *shell_script.Call) *Call {
	if c.Name != "git" {
		return nil
	}

	words := c.Words[1:]

	for len(words) > 0 {
		option, literal := shell_script.Argument(words[0])

		if !literal || !strings.HasPrefix(option, "-") {
			break
		}

		words = words[1:]

		if slices.Contains(constant.GitOptionsWithValue, option) &&
			len(words) > 0 {
			words = words[1:]
		}
	}

	if len(words) == 0 {
		return nil
	}

	subcommand, literal := shell_script.Argument(words[0])

	return New(subcommand, literal, words[1:], c.FromInput)
}

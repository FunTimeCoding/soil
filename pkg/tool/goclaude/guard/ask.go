package guard

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"slices"
	"strings"
)

func (c *GitCall) Ask() string {
	if !slices.Contains(constant.GitTreeWriters, c.subcommand) {
		return ""
	}

	var paths []string

	for _, w := range c.arguments {
		value, _ := literalWord(w)

		if !strings.HasPrefix(value, "-") {
			paths = append(paths, value)
		}
	}

	target := constant.GitWholeTree

	if len(paths) > 0 {
		target = strings.Join(paths, " ")
	}

	return fmt.Sprintf(constant.GitAskReason, c.subcommand, target)
}

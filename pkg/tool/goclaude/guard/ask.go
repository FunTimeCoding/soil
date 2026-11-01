package guard

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"slices"
	"strings"
)

func (c *GitCall) Ask() string {
	if !slices.Contains(constant.GitTreeWriters, c.subcommand) ||
		c.readsOnly() {
		return ""
	}

	name := c.subcommand
	arguments := c.arguments

	if v := c.stashVerb(); v != "" {
		name = join.Space(c.subcommand, v)
		arguments = arguments[1:]
	}

	var paths []string

	for _, w := range arguments {
		value, _ := literalWord(w)

		if !strings.HasPrefix(value, "-") {
			paths = append(paths, value)
		}
	}

	target := constant.GitWholeTree

	if len(paths) > 0 {
		target = strings.Join(paths, " ")
	}

	return fmt.Sprintf(constant.GitAskReason, name, target)
}

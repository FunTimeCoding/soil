package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func nestedScripts(call *syntax.CallExpr) ([]string, bool) {
	words := scriptWords(call)

	if len(words) == 0 {
		return nil, false
	}

	var parts []string

	for _, w := range words {
		value, readable := wordValue(w)

		if !readable {
			return nil, constant.GitMention.MatchString(printedWord(w))
		}

		parts = append(parts, value)
	}

	return []string{strings.Join(parts, " ")}, false
}

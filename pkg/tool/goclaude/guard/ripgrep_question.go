package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"mvdan.cc/sh/v3/syntax"
	"slices"
	"strings"
)

func ripgrepQuestion(command string) string {
	configured := strings.Contains(
		command,
		constant.RipgrepConfigurationVariable,
	)
	var reasons []string
	e := walkCalls(
		command,
		func(call *syntax.CallExpr) {
			r := ripgrepReason(call, configured)

			if r != "" && !slices.Contains(reasons, r) {
				reasons = append(reasons, r)
			}
		},
	)

	if e != nil && constant.RipgrepInvocation.MatchString(command) {
		return constant.RipgrepUnreadableReason
	}

	return strings.Join(reasons, "; ")
}

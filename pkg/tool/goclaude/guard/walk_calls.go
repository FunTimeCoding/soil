package guard

import (
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func walkCalls(
	command string,
	visit func(*syntax.CallExpr),
) error {
	file, e := syntax.NewParser().Parse(strings.NewReader(command), "")

	if e != nil {
		return e
	}

	var nested error
	syntax.Walk(
		file,
		func(node syntax.Node) bool {
			call, okay := node.(*syntax.CallExpr)

			if !okay || len(call.Args) == 0 {
				return true
			}

			visit(call)

			if words := scriptWords(call); len(words) > 0 {
				f := walkCalls(scriptText(words), visit)

				if f != nil && nested == nil {
					nested = f
				}
			}

			return true
		},
	)

	return nested
}

package reach

import (
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func ResolveCommand(
	cwd string,
	command string,
) *CommandWrites {
	result := &CommandWrites{}
	file, e := syntax.NewParser().Parse(strings.NewReader(command), "")

	if e != nil {
		result.Unresolved = true

		return result
	}

	directory := cwd
	syntax.Walk(
		file,
		func(n syntax.Node) bool {
			switch x := n.(type) {
			case *syntax.Stmt:
				addWords(result, directory, redirectTargets(x))
			case *syntax.CallExpr:
				directory = visitCall(x, directory, result)
			}

			return true
		},
	)

	return result
}

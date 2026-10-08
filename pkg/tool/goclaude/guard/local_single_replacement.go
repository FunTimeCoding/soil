package guard

import (
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func localSingleReplacement(command string) bool {
	file, e := syntax.NewParser().Parse(strings.NewReader(command), "")

	if e != nil {
		return false
	}

	var edits []*Edit
	var calls []*syntax.CallExpr
	syntax.Walk(
		file,
		func(node syntax.Node) bool {
			switch x := node.(type) {
			case *syntax.Stmt:
				source, extracted := pythonSource(x)

				if !extracted {
					return true
				}

				if target, single := singleReplacement(source); single {
					edits = append(edits, newEdit(x.End().Offset(), target))
				}
			case *syntax.CallExpr:
				if len(x.Args) > 0 {
					calls = append(calls, x)
				}
			}

			return true
		},
	)

	for _, d := range edits {
		if !restored(d, calls) {
			return true
		}
	}

	return false
}

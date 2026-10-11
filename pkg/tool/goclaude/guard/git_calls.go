package guard

import (
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func gitCalls(command string) ([]*GitCall, error) {
	file, e := syntax.NewParser().Parse(strings.NewReader(command), "")

	if e != nil {
		return nil, e
	}

	var result []*GitCall
	syntax.Walk(
		file,
		func(node syntax.Node) bool {
			call, okay := node.(*syntax.CallExpr)

			if !okay || len(call.Args) == 0 {
				return true
			}

			if c := findGitCall(call); c != nil {
				result = append(result, c)
			}

			return true
		},
	)

	return result, nil
}

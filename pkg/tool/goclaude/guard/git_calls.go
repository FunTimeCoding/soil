package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
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

			scripts, hidden := nestedScripts(call)

			if hidden {
				result = append(result, newHiddenGitCall())
			}

			for _, s := range scripts {
				inner, f := gitCalls(s)

				if f != nil && constant.GitMention.MatchString(s) {
					result = append(result, newHiddenGitCall())

					continue
				}

				result = append(result, inner...)
			}

			return true
		},
	)

	return result, nil
}

package shell_script

import (
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func (s *Script) walk(in *Script) {
	file, e := syntax.NewParser().Parse(strings.NewReader(in.Text), "")

	if e != nil {
		return
	}

	in.Parsed = true
	syntax.Walk(
		file,
		func(node syntax.Node) bool {
			statement, okay := node.(*syntax.Stmt)

			if !okay {
				return true
			}

			x, okay := statement.Cmd.(*syntax.CallExpr)

			if !okay || len(x.Args) == 0 {
				return true
			}

			var nested *Script

			if words := scriptWords(x); len(words) > 0 {
				nested = nestedScript(words, in)
			}

			s.Calls = append(s.Calls, newCall(statement, x, in, nested))

			if nested != nil {
				s.Scripts = append(s.Scripts, nested)
				s.walk(nested)
			}

			return true
		},
	)
}

package guard

import (
	"mvdan.cc/sh/v3/syntax"
	"path"
)

func pythonSource(s *syntax.Stmt) (string, bool) {
	call, okay := s.Cmd.(*syntax.CallExpr)

	if !okay || len(call.Args) == 0 {
		return "", false
	}

	name := path.Base(call.Args[0].Lit())

	if name != "python" && name != "python3" {
		return "", false
	}

	for i, argument := range call.Args[1:] {
		value := argument.Lit()

		if value == "-c" && i+2 < len(call.Args) {
			return wordValue(call.Args[i+2])
		}

		if value != "-" && (value == "" || value[0] != '-') {
			return "", false
		}
	}

	for _, r := range s.Redirs {
		if r.Hdoc != nil && (r.Op == syntax.Hdoc || r.Op == syntax.DashHdoc) {
			return wordValue(r.Hdoc)
		}
	}

	return "", false
}

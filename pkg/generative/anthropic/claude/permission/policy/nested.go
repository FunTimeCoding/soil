package policy

import "mvdan.cc/sh/v3/syntax"

func (w *Walk) nested(node syntax.Node) bool {
	result := true
	syntax.Walk(
		node,
		func(n syntax.Node) bool {
			switch x := n.(type) {
			case *syntax.CmdSubst:
				if !w.scoped(func() bool { return w.statements(x.Stmts) }) {
					result = false
				}

				return false
			case *syntax.ProcSubst:
				if !w.scoped(func() bool { return w.statements(x.Stmts) }) {
					result = false
				}

				return false
			}

			return true
		},
	)

	return result
}

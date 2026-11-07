package policy

import "mvdan.cc/sh/v3/syntax"

func (w *Walk) match(x *syntax.CaseClause) bool {
	result := w.nested(x.Word)

	for _, item := range x.Items {
		for _, p := range item.Patterns {
			if !w.nested(p) {
				result = false
			}
		}

		if !w.conditional(func() bool { return w.statements(item.Stmts) }) {
			result = false
		}
	}

	return result
}

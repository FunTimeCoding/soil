package policy

import "mvdan.cc/sh/v3/syntax"

func (w *Walk) branch(x *syntax.IfClause) bool {
	condition := w.statements(x.Cond)
	then := w.conditional(func() bool { return w.statements(x.Then) })
	otherwise := x.Else == nil ||
		w.conditional(func() bool { return w.branch(x.Else) })

	return condition && then && otherwise
}

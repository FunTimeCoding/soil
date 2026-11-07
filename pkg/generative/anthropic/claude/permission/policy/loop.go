package policy

import "mvdan.cc/sh/v3/syntax"

func (w *Walk) loop(x *syntax.ForClause) bool {
	words := w.nested(x.Loop)

	if i, okay := x.Loop.(*syntax.WordIter); okay && i.Name != nil {
		delete(w.variables, i.Name.Value)
	}

	return w.conditional(func() bool { return w.statements(x.Do) }) && words
}

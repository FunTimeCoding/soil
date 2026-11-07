package policy

import "mvdan.cc/sh/v3/syntax"

func (w *Walk) binary(x *syntax.BinaryCmd) bool {
	switch x.Op {
	case syntax.Pipe, syntax.PipeAll:
		left := w.scoped(func() bool { return w.statement(x.X) })

		return w.conditional(func() bool { return w.statement(x.Y) }) && left
	case syntax.OrStmt:
		left := w.statement(x.X)

		return w.conditional(func() bool { return w.statement(x.Y) }) && left
	default:
		left := w.statement(x.X)

		return w.statement(x.Y) && left
	}
}

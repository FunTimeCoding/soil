package policy

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"mvdan.cc/sh/v3/syntax"
)

func (w *Walk) command(c syntax.Command) bool {
	switch x := c.(type) {
	case nil:
		return true
	case *syntax.CallExpr:
		return w.call(x)
	case *syntax.BinaryCmd:
		return w.binary(x)
	case *syntax.Subshell:
		return w.scoped(func() bool { return w.statements(x.Stmts) })
	case *syntax.Block:
		return w.statements(x.Stmts)
	case *syntax.ForClause:
		return w.loop(x)
	case *syntax.WhileClause:
		return w.conditional(
			func() bool {
				condition := w.statements(x.Cond)

				return w.statements(x.Do) && condition
			},
		)
	case *syntax.IfClause:
		return w.branch(x)
	case *syntax.CaseClause:
		return w.match(x)
	case *syntax.TestClause:
		return w.nested(x)
	case *syntax.ArithmCmd:
		return w.nested(x)
	case *syntax.DeclClause:
		return w.declare(x)
	case *syntax.TimeClause:
		return x.Stmt == nil || w.statement(x.Stmt)
	default:
		return w.block(constant.BlockerConstruct)
	}
}

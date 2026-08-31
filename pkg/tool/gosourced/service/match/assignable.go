package match

import (
	"go/ast"
	"go/types"
)

func (m *Matcher) assignable(expression ast.Expr, typeText string) bool {
	actual := m.information.TypeOf(expression)

	if actual == nil {
		return false
	}

	target, e := types.Eval(m.set, m.scope, expression.Pos(), typeText)

	if e != nil || target.Type == nil {
		return false
	}

	return types.AssignableTo(actual, target.Type)
}

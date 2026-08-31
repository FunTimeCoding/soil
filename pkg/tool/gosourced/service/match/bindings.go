package match

import "go/ast"

func (m *Matcher) Bindings() map[string]ast.Expr {
	return m.bindings
}
